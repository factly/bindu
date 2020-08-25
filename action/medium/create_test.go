package medium

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestMediumCreate(t *testing.T) {

	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount(url, Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable medium", func(t *testing.T) {

		e.POST(url).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create medium", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, organisation_id FROM "bi_medium"`)).
			WithArgs(fmt.Sprint(data["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_medium"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, data["name"], data["slug"], data["type"], byteData, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], data["slug"], data["type"], byteData))

		e.POST(url).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(data)
		mock.ExpectationsWereMet()

	})

	t.Run("create medium with slug is empty", func(t *testing.T) {

		slug := slug.Make(fmt.Sprint(mediumWithoutSlug["name"]))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT slug, organisation_id FROM "bi_medium"`)).
			WithArgs(slug+"%", 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "bi_medium"`)).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, mediumWithoutSlug["name"], slug, mediumWithoutSlug["type"], byteData, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug", "type", "url"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, mediumWithoutSlug["name"], slug, mediumWithoutSlug["type"], byteData))
		resObj := map[string]interface{}{
			"name": "Politics",
			"slug": slug,
		}

		e.POST(url).
			WithHeaders(headers).
			WithJSON(mediumWithoutSlug).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(resObj)

		mock.ExpectationsWereMet()
	})

}
