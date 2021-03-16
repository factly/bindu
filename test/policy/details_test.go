package policy

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

func TestDetails(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Get detail success", func(t *testing.T) {
		test.CheckSpace(mock)

		// Splits string of ID to retrieve the name of the policy. The name is in the last index, hence the split
		var id = strings.Split(test.Dummy_SingleMock["id"].(string), ":")

		e.GET(path).
			WithPath("policy_id", 1).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("name").Equal(id[len(id)-1])

	})

}
