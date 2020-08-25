package medium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestMediumUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount(url, Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid medium id", func(t *testing.T) {
		e.PUT(urlWithPath).
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("medium record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(columns))

		e.PUT(urlWithPath).
			WithPath("medium_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update medium", func(t *testing.T) {
		updatedMedium := map[string]interface{}{
			"name": "Elections",
			"slug": "testing",
			"type": "jpg",
			"url": `{"image": { 
				"src": "Images/election/Sun.png",
				"name": "sun1",
				"hOffset": 250,
				"vOffset": 250,
				"alignment": "center"
			}}`,
		}

		updatedByteData, _ := json.Marshal(updatedMedium["url"])

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, "Elections", "testing", "png", updatedByteData))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_medium\" SET (.+)  WHERE (.+) \"bi_medium\".\"id\" = `).
			WithArgs(updatedMedium["name"], updatedMedium["slug"], updatedMedium["type"], test.AnyTime{}, updatedByteData, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, updatedMedium["name"], updatedMedium["slug"], updatedMedium["type"], updatedByteData))

		e.PUT(urlWithPath).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedMedium)

	})

	t.Run("update medium by id with empty slug", func(t *testing.T) {
		updatedMedium := map[string]interface{}{
			"name": "Sun",
			"slug": "",
			"type": "jpg",
			"url": `{"image": {
				"src": "Images/Sun.png",
				"name": "sun1",
				"hOffset": 250,
				"vOffset": 250,
				"alignment": "right"
			}}`,
		}

		updatedByteData, _ := json.Marshal(updatedMedium["url"])
		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, "Sun", "sun", "png", byteData))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_medium"`).
			WithArgs("sun%", 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, "Sun", "sun", "png", byteData))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_medium\" SET (.+)  WHERE (.+) \"bi_medium\".\"id\" = `).
			WithArgs(updatedMedium["name"], "sun-1", updatedMedium["type"], test.AnyTime{}, updatedByteData, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, updatedMedium["name"], "sun-1", updatedMedium["type"], updatedByteData))

		resObj := map[string]interface{}{
			"name": "Sun",
			"slug": "sun-1",
			"type": "jpg",
			"url": `{"image": {
				"src": "Images/Sun.png",
				"name": "sun1",
				"hOffset": 250,
				"vOffset": 250,
				"alignment": "right"
			}}`,
		}

		e.PUT(urlWithPath).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(resObj)

	})

	t.Run("update medium with different slug", func(t *testing.T) {

		updatedMedium := map[string]interface{}{
			"name": "Graph",
			"slug": "testing-slug",
			"type": "jpg",
			"url": `{"image": { 
				"src": "Images/graphs/Bar.png",
				"name": "Graph",
				"hOffset": 250,
				"vOffset": 250,
				"alignment": "center"
			}}`,
		}
		updatedByteData, _ := json.Marshal(updatedMedium["url"])

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, "Graph", "testing", "png", updatedByteData))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_medium"`).
			WithArgs(fmt.Sprint(updatedMedium["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows(columns))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_medium\" SET (.+)  WHERE (.+) \"bi_medium\".\"id\" = `).
			WithArgs(updatedMedium["name"], "testing-slug", updatedMedium["type"], test.AnyTime{}, updatedByteData, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, updatedMedium["name"], "testing-slug", updatedMedium["type"], updatedByteData))

		e.PUT(urlWithPath).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedMedium)

	})

}
