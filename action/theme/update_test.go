package theme

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestThemeUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/themes", Router())

	ts := httptest.NewServer(r)
	gock.New(ts.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer ts.Close()

	theme := &model.Theme{
		Name:           "Theme",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)

	t.Run("invalid theme id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/themes/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("theme record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/themes/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("update theme", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		var jsonStr = []byte(`
		{
			"name": "Theme sample"
		}`)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/themes/", theme.Base.ID), bytes.NewBuffer(jsonStr), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Theme sample" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Theme sample")
		}

	})

}
