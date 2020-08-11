package chart

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestChartCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Post("/charts", create)

	ts := httptest.NewServer(r)
	defer ts.Close()

	category := &model.Category{
		Name:           "Sports",
		OrganisationID: 1,
	}

	tag := &model.Tag{
		Name:           "Agriculture",
		OrganisationID: 1,
	}

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)
	config.DB.Model(&model.Tag{}).Create(&tag)
	config.DB.Model(&model.Category{}).Create(&category)

	t.Run("Unprocessable chart", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/charts", nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("create chart", func(t *testing.T) {
		reqBody := chart{
			Title:   "Pie Chart",
			Slug:    "pie_chart",
			ThemeID: theme.Base.ID,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}

		requestByte, _ := json.Marshal(reqBody)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/charts", bytes.NewBuffer(requestByte), headers)

		respBody := (resp).(map[string]interface{})

		t.Log(respBody)

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["title"] != "Pie Chart" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["title"], "Pie Chart")
		}

	})

	t.Run("Invalid theme id, medium id, category ids & tag ids", func(t *testing.T) {

		reqBody := chart{
			Title:            "Bar",
			ThemeID:          100,
			FeaturedMediumID: 100,
			TagIDs: []uint{
				100,
			},
			CategoryIDs: []uint{
				100,
			},
		}

		requestByte, _ := json.Marshal(reqBody)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "POST", "/charts", bytes.NewBuffer(requestByte), headers)

		if statusCode != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusInternalServerError)
		}

	})

	t.Run("create chart with slug is empty", func(t *testing.T) {

		reqBody := chart{
			Title:   "Bar",
			Slug:    "bar",
			ThemeID: theme.Base.ID,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}

		requestByte, _ := json.Marshal(reqBody)
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "POST", "/charts", bytes.NewBuffer(requestByte), headers)

		respBody := (resp).(map[string]interface{})

		if statusCode != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusCreated)
		}

		if respBody["slug"] != "bar" {
			t.Errorf("handler returned wrong title: got %v want %v", respBody["slug"], "bar")
		}

	})

}
