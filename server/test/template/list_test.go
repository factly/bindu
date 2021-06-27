package template

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/category"
	"github.com/factly/bindu-server/test/medium"
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

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns))

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array().
			Empty()

		test.ExpectationsMet(t, mock)
	})

	t.Run("get list of templates", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow("1", time.Now(), time.Now(), nil, 1, 1, templateList[0]["title"], templateList[0]["slug"], templateList[0]["spec"], templateList[0]["properties"], templateList[0]["category_id"], templateList[0]["medium_id"], templateList[0]["is_default"], templateList[0]["mode"], templateList[0]["description"], templateList[0]["html_description"], 1).
				AddRow("2", time.Now(), time.Now(), nil, 1, 1, templateList[1]["title"], templateList[1]["slug"], templateList[1]["spec"], templateList[1]["properties"], templateList[1]["category_id"], templateList[1]["medium_id"], templateList[1]["is_default"], templateList[1]["mode"], templateList[1]["description"], templateList[1]["html_description"], 1))

		category.SelectMock(mock)
		medium.SelectMock(mock)

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array()
		test.ExpectationsMet(t, mock)
	})

}
