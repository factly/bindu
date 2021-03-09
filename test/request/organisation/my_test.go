package organisation

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

func TestMyOrganisationRequestList(t *testing.T) {
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

	t.Run("get empty list of requests for organisation", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(Columns))

		e.GET(myPath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 0})

		test.ExpectationsMet(t, mock)
	})

	t.Run("get list of requests for organisation", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(len(requestList)))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(Columns).
				AddRow(1, time.Now(), time.Now(), nil, 1, 1, requestList[0]["title"], requestList[0]["description"], requestList[0]["status"], requestList[0]["organisation_id"], requestList[0]["spaces"]).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, requestList[1]["title"], requestList[1]["description"], requestList[1]["status"], requestList[1]["organisation_id"], requestList[1]["spaces"]))

		e.GET(myPath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": len(requestList)}).
			Value("nodes").
			Array().Element(0).Object().ContainsMap(requestList[0])

		test.ExpectationsMet(t, mock)
	})

	t.Run("get paiganation list of requests for organisation", func(t *testing.T) {
		test.CheckSpace(mock)

		mock.ExpectQuery(countQuery).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		mock.ExpectQuery(selectQuery).
			WillReturnRows(sqlmock.NewRows(Columns).
				AddRow(2, time.Now(), time.Now(), nil, 1, 1, requestList[1]["title"], requestList[1]["description"], requestList[1]["status"], requestList[1]["organisation_id"], requestList[1]["spaces"]))

		e.GET(basePath).
			WithHeaders(headers).
			WithQueryObject(map[string]interface{}{
				"page":  2,
				"limit": 1,
			}).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 1}).
			Value("nodes").
			Array().Element(0).Object().ContainsMap(requestList[1])

		test.ExpectationsMet(t, mock)

	})
}
