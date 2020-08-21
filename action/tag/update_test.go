package tag

import (
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

func TestTagUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount(url, Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	var updatedTag = map[string]interface{}{
		"name": "Politics",
		"slug": "politics",
	}

	t.Run("invalid tag id", func(t *testing.T) {
		e.PUT(urlWithPath).
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("tag record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(tagProps))

		e.PUT(urlWithPath).
			WithPath("tag_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update tag", func(t *testing.T) {

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, "Elections", "politics"))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_tag\" SET (.+)  WHERE (.+) \"bi_tag\".\"id\" = `).
			WithArgs(updatedTag["name"], updatedTag["slug"], test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, updatedTag["name"], updatedTag["slug"]))

		e.PUT(urlWithPath).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTag).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTag)

	})

	t.Run("update tag by id with empty slug", func(t *testing.T) {

		updatedTag := map[string]interface{}{
			"name": "Politics",
			"slug": "",
		}
		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, "Politics", "politics"))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_tag"`).
			WithArgs("politics%", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, "Politics", "politics"))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_tag\" SET (.+)  WHERE (.+) \"bi_tag\".\"id\" = `).
			WithArgs(updatedTag["name"], "politics-1", test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, updatedTag["name"], "politics-1"))

		resObj := map[string]interface{}{
			"name": "Politics",
			"slug": "politics-1",
		}

		e.PUT(urlWithPath).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTag).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(resObj)

	})

	t.Run("update tag with different slug", func(t *testing.T) {
		updatedTag := map[string]interface{}{
			"name": "Politics",
			"slug": "testing-slug",
		}
		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, "Politics", "slug"))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_tag"`).
			WithArgs(fmt.Sprint(updatedTag["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE \"bi_tag\" SET (.+)  WHERE (.+) \"bi_tag\".\"id\" = `).
			WithArgs(updatedTag["name"], "testing-slug", test.AnyTime{}, 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, updatedTag["name"], "testing-slug"))

		e.PUT(urlWithPath).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTag).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTag)

	})

}
