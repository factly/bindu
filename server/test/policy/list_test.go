package policy

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gock.v1"
)

func TestList(t *testing.T) {
	mock := test.SetupMockDB()

	test.MockServers()
	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("List success", func(t *testing.T) {
		test.CheckSpace(mock)

		// Splits string of ID to retrieve the name of the policy. The name is in the last index, hence the split
		var text = strings.Split(test.Dummy_KetoPolicy[1]["id"].(string), ":")

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("nodes").Array().Element(0).Object().Value("name").Equal(text[len(text)-1])
	})

	t.Run("when keto cannot fetch policies", func(t *testing.T) {
		test.DisableKetoGock(testServer.URL)
		gock.New(viper.GetString("keto_url")).
			Post("/engines/acp/ory/regex/allowed").
			Persist().
			Reply(http.StatusOK)

		test.CheckSpace(mock)

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusInternalServerError)
	})

}
