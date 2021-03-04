package tag

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

func TestTagUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid tag id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
	})

	t.Run("cannot decode tag", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})
	t.Run("Unprocessable tag", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("tag record not found", func(t *testing.T) {
		test.CheckSpace(mock)
		recordNotFoundMock(mock)

		e.PUT(path).
			WithPath("tag_id", "100").
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update tag", func(t *testing.T) {
		test.CheckSpace(mock)

		updatedTag := map[string]interface{}{
			"name": "Elections",
			"slug": "elections",
		}

		SelectMock(mock)

		tagUpdateMock(mock, updatedTag)

		selectAfterUpdate(mock, updatedTag)

		e.PUT(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTag).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTag)

	})

	t.Run("update tag by id with empty slug", func(t *testing.T) {

		test.CheckSpace(mock)
		updatedTag := map[string]interface{}{
			"name": "Elections",
			"slug": "elections-1",
		}
		SelectMock(mock)

		mock.ExpectQuery(`SELECT slug, space_id FROM "bi_tag"`).
			WithArgs("elections%", 1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, updatedTag["name"], "elections", 1))

		tagUpdateMock(mock, updatedTag)

		selectAfterUpdate(mock, updatedTag)

		e.PUT(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(dataWithoutSlug).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTag)

	})

	t.Run("update tag with different slug", func(t *testing.T) {
		test.CheckSpace(mock)
		updatedTag := map[string]interface{}{
			"name": "Elections",
			"slug": "testing-slug",
		}
		SelectMock(mock)

		mock.ExpectQuery(`SELECT slug, space_id FROM "bi_tag"`).
			WithArgs(fmt.Sprint(updatedTag["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "space_id"}))

		tagUpdateMock(mock, updatedTag)

		selectAfterUpdate(mock, updatedTag)

		e.PUT(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			WithJSON(updatedTag).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedTag)

	})

}
