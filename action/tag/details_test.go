package tag

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

func TestTagDetails(t *testing.T) {
	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid tag id", func(t *testing.T) {
		e.GET("/tags/{tag_id}").
			WithPath("tag_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("tag record not found", func(t *testing.T) {
		mock.ExpectQuery(selectQuery).
			WithArgs(100, 1).
			WillReturnRows(sqlmock.NewRows(tagProps))

		e.GET("/tags/{tag_id}").
			WithPath("tag_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("get tag by id", func(t *testing.T) {

		mock.ExpectQuery(selectQuery).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows(tagProps).
				AddRow(1, time.Now(), time.Now(), nil, data["name"], data["slug"]))

		e.GET("/tags/{tag_id}").
			WithPath("tag_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().ContainsMap(data)
	})

}
