package template

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestTemplateList(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get empty list of templates", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 0})

		test.ExpectationsMet(t, mock)
	})

	t.Run("get list of templates", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(templateList)))

		mock.ExpectQuery(selectQuery).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, templateList[0]["title"], templateList[0]["slug"], templateList[0]["schema"], templateList[0]["properties"], templateList[0]["medium_id"], 1).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, templateList[1]["title"], templateList[1]["slug"], templateList[1]["schema"], templateList[1]["properties"], templateList[1]["medium_id"], 1))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(templateList)}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(templateList[0])

		test.ExpectationsMet(t, mock)
	})

}
