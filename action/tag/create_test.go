package tag

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestTagCreate(t *testing.T) {

	mock := test.SetupMockDB()
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount(url, Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable tag", func(t *testing.T) {

		e.POST(url).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create tag", func(t *testing.T) {

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_tag"`).
			WithArgs(fmt.Sprint(data["slug"], "%"), 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_tag"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, data["name"], data["slug"], "", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, data["name"], data["slug"]))

		e.POST(url).
			WithHeaders(headers).
			WithJSON(data).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(data)
		mock.ExpectationsWereMet()

	})

	t.Run("create tag with slug is empty", func(t *testing.T) {

		slug := slug.Make(fmt.Sprint(tagWithoutSlug["name"]))

		mock.ExpectQuery(`SELECT slug, organisation_id FROM "bi_tag"`).
			WithArgs(slug+"%", 1).
			WillReturnRows(sqlmock.NewRows([]string{"slug", "organisation_id"}))
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_tag"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, tagWithoutSlug["name"], slug, "", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "organisation_id", "name", "slug"}).
				AddRow(1, time.Now(), time.Now(), nil, 1, tagWithoutSlug["name"], slug))
		resObj := map[string]interface{}{
			"name": "Politics",
			"slug": slug,
		}

		e.POST(url).
			WithHeaders(headers).
			WithJSON(tagWithoutSlug).
			Expect().
			Status(http.StatusCreated).JSON().Object().ContainsMap(resObj)

		mock.ExpectationsWereMet()
	})

}
