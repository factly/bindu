package chart

import (
	"bytes"
	"encoding/json"
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

func TestChartUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/charts", Router())

	ts := httptest.NewServer(r)
	gock.New(ts.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer ts.Close()

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	category := &model.Category{
		Name:           "Sports",
		OrganisationID: 1,
	}

	tag := &model.Tag{
		Name:           "Agriculture",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Tag{}).Create(&tag)
	config.DB.Model(&model.Category{}).Create(&category)
	config.DB.Model(&model.Theme{}).Create(&theme)

	result := &model.Chart{
		Title:          "Test",
		ThemeID:        theme.Base.ID,
		OrganisationID: 1,
	}

	config.DB.Model(&model.Chart{}).Create(&result)

	t.Run("invalid chart id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/charts/invalid_id", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("chart record not found", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "PUT", "/charts/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("update chart", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		reqBody := &chart{
			Title:          "Maps",
			Slug:           "maps",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}

		requestByte, _ := json.Marshal(reqBody)
		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/charts/", result.Base.ID), bytes.NewBuffer(requestByte), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["title"] != "Maps" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["title"], "Maps")
		}

	})

	t.Run("update chart by id with empty slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		reqBody := &model.Chart{
			Title: "Maps",
			Slug:  "",
		}

		requestByte, _ := json.Marshal(reqBody)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/charts/", result.Base.ID), bytes.NewBuffer(requestByte), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["title"] != "Maps" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["title"], "Maps")
		}

	})

	t.Run("update chart with different slug", func(t *testing.T) {

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}

		reqBody := &model.Chart{
			Title: "Maps",
			Slug:  "map-sample",
		}

		requestByte, _ := json.Marshal(reqBody)

		resp, statusCode := test.Request(t, ts, "PUT", fmt.Sprint("/charts/", result.Base.ID), bytes.NewBuffer(requestByte), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if respBody["title"] != "Maps" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["title"], "Maps")
		}

	})

}
