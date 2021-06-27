package space

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/space"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gock.v1"
)

func TestSpaceRequestCreate(t *testing.T) {

	mock := test.SetupMockDB()
	test.SetEnv()

	test.MockServers()
	defer gock.DisableNetworking()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable request body", func(t *testing.T) {
		e.POST(basePath).
			WithJSON(invalidData).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("Undecodable request body", func(t *testing.T) {
		e.POST(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
		test.ExpectationsMet(t, mock)
	})

	t.Run("Space for the request not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "bi_space"`)).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(space.Columns))

		e.POST(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
		test.ExpectationsMet(t, mock)
	})

	t.Run("User not owner of the organisation", func(t *testing.T) {
		space.SelectQuery(mock, 1)

		test.DisableKavachGock(testServer.URL)

		gock.New(viper.GetString("kavach_url") + "/organisations/my").
			Persist().
			Reply(http.StatusOK).
			JSON([]map[string]interface{}{
				map[string]interface{}{
					"id":         1,
					"created_at": time.Now(),
					"updated_at": time.Now(),
					"deleted_at": nil,
					"title":      "test org",
					"slug":       "test-org",
					"permission": map[string]interface{}{
						"id":              1,
						"created_at":      time.Now(),
						"updated_at":      time.Now(),
						"deleted_at":      nil,
						"user_id":         1,
						"user":            nil,
						"organisation_id": 1,
						"organisation":    nil,
						"role":            "member",
					},
				},
			})

		e.POST(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnauthorized)
		test.ExpectationsMet(t, mock)

	})

	t.Run("Create space request", func(t *testing.T) {
		gock.Off()
		test.MockServers()
		gock.New(testServer.URL).EnableNetworking().Persist()
		defer gock.DisableNetworking()

		space.SelectQuery(mock, 1)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "bi_space_permission_request"`).
			WithArgs(test.AnyTime{}, test.AnyTime{}, nil, 1, 1, Data["title"], Data["description"], "pending", Data["charts"], Data["space_id"]).
			WillReturnRows(sqlmock.
				NewRows([]string{"id"}).
				AddRow(1))
		mock.ExpectCommit()

		e.POST(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusCreated)
		test.ExpectationsMet(t, mock)
	})

}
