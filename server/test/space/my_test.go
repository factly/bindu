package space

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/test/medium"
	"github.com/factly/bindu-server/util/test"
	"github.com/gavv/httpexpect"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gock.v1"
)

func TestSpaceMy(t *testing.T) {
	mock := test.SetupMockDB()

	defer gock.Disable()
	test.MockServers()
	defer gock.DisableNetworking()

	testServer := httptest.NewServer(action.RegisterRoutes())
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("get my spaces", func(t *testing.T) {
		SelectQuery(mock, 1)

		medium.SelectMock(mock)
		medium.SelectMock(mock)
		medium.SelectMock(mock)
		medium.SelectMock(mock)

		e.GET(basePath).
			WithHeader("X-User", "1").
			Expect().
			Status(http.StatusOK).
			JSON().
			Array().
			Element(0).
			Object().
			Value("spaces").
			Array().
			Element(0).
			Object().
			ContainsMap(resData)

		test.ExpectationsMet(t, mock)
	})

	t.Run("invalid space header", func(t *testing.T) {
		e.GET(basePath).
			WithHeader("X-User", "invalid").
			Expect().
			Status(http.StatusUnauthorized)
	})

	t.Run("when kavach is down", func(t *testing.T) {
		test.DisableKavachGock(testServer.URL)

		e.GET(basePath).
			WithHeader("X-User", "1").
			Expect().
			Status(http.StatusInternalServerError)
	})

	t.Run("when member requests his spaces", func(t *testing.T) {
		test.DisableKavachGock(testServer.URL)
		SelectQuery(mock, 1)

		medium.SelectMock(mock)
		medium.SelectMock(mock)
		medium.SelectMock(mock)
		medium.SelectMock(mock)

		gock.New(viper.GetString("kavach_url") + "/organisations/my").
			Persist().
			Reply(http.StatusOK).
			JSON(test.Dummy_Org_Member_List)

		e.GET(basePath).
			WithHeader("X-User", "1").
			Expect().
			Status(http.StatusOK)
	})

}
