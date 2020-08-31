package tag

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestTagDetails(t *testing.T) {
	mock := test.SetupMockDB()

	testServer := httptest.NewServer(Routes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid tag id", func(t *testing.T) {
		e.GET(path).
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("tag record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100).
			WillReturnRows(sqlmock.NewRows(columns))

		e.GET(path).
			WithPath("tag_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("get tag by id", func(t *testing.T) {
		tagSelectWithOutOrg(mock)

		e.GET(path).
			WithPath("tag_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(data)
	})

}
