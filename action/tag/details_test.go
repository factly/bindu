package tag

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
)

func TestTagDetails(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Get("/tags/{tag_id}", details)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid tag id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/tags/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("tag record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "GET", "/tags/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("get tag by id", func(t *testing.T) {
		tag := &model.Tag{
			Name:           "Agriculture",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Tag{}).Create(&tag)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		resp, statusCode := test.Request(t, ts, "GET", fmt.Sprint("/tags/", tag.Base.ID), nil, headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["name"] != "Agriculture" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["name"], "Agriculture")
		}

	})

}
