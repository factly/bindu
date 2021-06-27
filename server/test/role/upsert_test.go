package role

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"gopkg.in/h2non/gock.v1"
)

func TestUpsert(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	// Create a role
	t.Run("Successful create role", func(t *testing.T) {

		test.CheckSpace(mock)

		e.PUT(basePath).
			WithJSON(Data).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().Value("name").Equal(Data["name"])
	})

	t.Run("undecodable role body", func(t *testing.T) {
		test.CheckSpace(mock)

		e.PUT(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

}
