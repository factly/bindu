package policy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestUpsertPolicy(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	// Create a policy
	t.Run("Successful create policy", func(t *testing.T) {

		test.CheckSpace(mock)

		e.PUT(basePath).
			WithJSON(policy_test).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).JSON().Object().Value("name").Equal(policy_test["name"])
	})

	t.Run("undecodable policy body", func(t *testing.T) {
		test.CheckSpace(mock)

		e.PUT(basePath).
			WithJSON(undecodable_policy).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

}
