package medium

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

var updateData = map[string]interface{}{
	"name": "Politics",
	"type": "jpg",
	"url": `{"image": { 
		"src": "Images/politics/Sun.png",
		"name": "sun1",
		"hOffset": 250,
		"vOffset": 250,
		"alignment": "center"
	}}`,
}

func TestMediumUpdate(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)
	res := map[string]interface{}{
		"name": "Politics",
		"type": "jpg",
		"url": `{"image": { 
			"src": "Images/politics/Sun.png",
			"name": "sun1",
			"hOffset": 250,
			"vOffset": 250,
			"alignment": "center"
		}}`,
	}

	t.Run("invalid medium id", func(t *testing.T) {
		e.PUT(path).
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})
	t.Run("cannot decode medium", func(t *testing.T) {

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable medium", func(t *testing.T) {

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("medium record not found", func(t *testing.T) {
		recordNotFoundMock(mock)

		e.PUT(path).
			WithPath("medium_id", "100").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})
	t.Run("update medium", func(t *testing.T) {
		updatedMedium := updateData
		updatedMedium["slug"] = "politics"

		mediumSelectMock(mock)

		mediumUpdateMock(mock, updatedMedium)
		res["slug"] = "politics"

		selectAfterUpdate(mock, res)

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

	t.Run("update medium by id with empty slug", func(t *testing.T) {
		updatedMedium := updateData
		updatedMedium["slug"] = "politics-1"

		mediumSelectMock(mock)

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_medium"`).
			WithArgs("politics%", 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, updatedMedium["name"], "politics", updatedMedium["type"], byteData))

		mediumUpdateMock(mock, updatedMedium)

		res["slug"] = "politics-1"

		selectAfterUpdate(mock, res)

		updatedMedium["slug"] = ""

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

	t.Run("update medium with different slug", func(t *testing.T) {

		updatedMedium := updateData
		updatedMedium["slug"] = "politics-test"

		mediumSelectMock(mock)

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_medium"`).
			WithArgs(fmt.Sprint(updatedMedium["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows(columns))

		mediumUpdateMock(mock, updatedMedium)

		res["slug"] = "politics-test"
		selectAfterUpdate(mock, res)

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(res)

	})

}
