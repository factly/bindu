package medium

import (
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

func TestMediumDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/media", Router())

	ts := httptest.NewServer(r)
	gock.New(ts.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer ts.Close()

	t.Run("invalid medium id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/media/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("medium record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/media/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("get medium by id", func(t *testing.T) {
		medium := &model.Medium{
			Name:           "Medium sample",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Medium{}).Create(&medium)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		resp, statusCode := test.Request(t, ts, "GET", fmt.Sprint("/media/", medium.Base.ID), nil, headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Medium sample" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Medium sample")
		}

	})

}
