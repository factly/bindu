package medium

import (
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

func TestMediumDetails(t *testing.T) {
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
		e.GET(urlWithPath).
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("medium record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(mediumProps))

		e.GET(urlWithPath).
			WithPath("medium_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("get medium by id", func(t *testing.T) {

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(mediumProps).
				AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], data["slug"], data["type"], byteData))

		e.GET(urlWithPath).
			WithPath("medium_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(data)
	})

}
