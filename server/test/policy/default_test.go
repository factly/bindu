package policy

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/action/policy"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateDefaultPolicy(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	policy.DataFile = "../../data/policies.json"

	// Create a policy
	t.Run("create default policies", func(t *testing.T) {
		test.CheckSpace(mock)
		e.POST(defaultsPath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusCreated).
			JSON().
			Object().
			Value("nodes").
			Array()
		test.ExpectationsMet(t, mock)
	})

	t.Run("when cannot open data file", func(t *testing.T) {
		policy.DataFile = "nofile.json"
		test.CheckSpace(mock)

		e.POST(defaultsPath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusInternalServerError)
		test.ExpectationsMet(t, mock)
		policy.DataFile = "../../../../data/policies.json"
	})

	t.Run("when cannot parse data file", func(t *testing.T) {
		policy.DataFile = "invalidData.json"
		test.CheckSpace(mock)

		e.POST(defaultsPath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusInternalServerError)
		test.ExpectationsMet(t, mock)
		policy.DataFile = "../../../../data/policies.json"
	})

}
