package theme

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestThemeCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/themes", Router())

	themeOne := model.Theme{
		Name: "Light",
	}

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("Unprocessable theme", func(t *testing.T) {
		e.POST("/themes").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create theme", func(t *testing.T) {
		resObj := e.POST("/themes").
			WithHeaders(headers).
			WithJSON(themeOne).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("name").String().Equal("Light")

	})

}
