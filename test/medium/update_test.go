package medium

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/h2non/gock.v1"
)

var updateData = map[string]interface{}{
	"name":        "Image",
	"slug":        "image",
	"type":        "jpg",
	"title":       "Sample image",
	"description": "desc",
	"caption":     "sample",
	"alt_text":    "sample",
	"file_size":   100,
	"url": postgres.Jsonb{
		RawMessage: []byte(`{"raw":"http://testimage.com/test.jpg"}`),
	},
	"dimensions": "testdims",
}

func TestMediumUpdate(t *testing.T) {
	mock := test.SetupMockDB()
	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid medium id", func(t *testing.T) {
		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusBadRequest)
	})
	t.Run("cannot decode medium", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("Unprocessable medium", func(t *testing.T) {

		test.CheckSpace(mock)
		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(invalidData).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("medium record not found", func(t *testing.T) {
		test.CheckSpace(mock)
		recordNotFoundMock(mock)

		e.PUT(path).
			WithPath("medium_id", "100").
			WithJSON(data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})
	t.Run("update medium", func(t *testing.T) {
		test.CheckSpace(mock)
		updatedMedium := updateData
		updatedMedium["slug"] = "image"

		mediumSelectMock(mock)

		mediumUpdateMock(mock, updatedMedium)

		selectAfterUpdate(mock, updateData)

		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updatedMedium).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updateData)

	})

	t.Run("update medium by id with empty slug", func(t *testing.T) {
		updatedMedium := updateData
		updatedMedium["slug"] = "image"

		test.CheckSpace(mock)
		mediumSelectMock(mock)

		mock.ExpectQuery(`SELECT slug, space_id FROM "bi_medium"`).
			WithArgs("image%", 1).
			WillReturnRows(sqlmock.NewRows(columns))
		mediumUpdateMock(mock, updatedMedium)

		selectAfterUpdate(mock, updatedMedium)

		data["slug"] = ""
		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updatedMedium)
		data["slug"] = "image"
	})

	t.Run("update medium with different slug", func(t *testing.T) {
		test.CheckSpace(mock)

		updatedMedium := updateData
		updatedMedium["slug"] = "image-test"

		mediumSelectMock(mock)

		mock.ExpectQuery(`SELECT slug, space_id FROM "bi_medium"`).
			WithArgs(fmt.Sprint(updatedMedium["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows(columns))

		mediumUpdateMock(mock, updatedMedium)

		selectAfterUpdate(mock, updatedMedium)

		updateData["slug"] = "image-test"
		e.PUT(path).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			WithJSON(updateData).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(updateData)
		updateData["slug"] = "image"
	})

}
