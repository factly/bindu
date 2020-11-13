package organisation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect/v2"
	"gopkg.in/h2non/gock.v1"
)

var basePath = "/organisations"

var orgList = []map[string]interface{}{{
	"id":         1,
	"deleted_at": nil,
	"title":      "test org",
	"slug":       "tesing",
	"permission": map[string]interface{}{
		"id":         1,
		"deleted_at": nil,
		"role":       "owner",
	}},
}

var headers = map[string]string{
	"X-Organisation": "1",
	"X-User":         "1",
}

func TestOrganisationList(t *testing.T) {
	test.SetEnv()
	defer gock.Disable()
	test.MockServer()
	defer gock.DisableNetworking()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get empty list of organisations", func(t *testing.T) {

		e.GET(basePath).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK).
			JSON().
			Object().
			ContainsMap(map[string]interface{}{"total": 1}).
			Value("nodes").
			Array().
			Element(0).
			Object().
			ContainsMap(orgList[0])

	})
}
