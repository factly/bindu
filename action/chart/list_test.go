package chart

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/test"
	"github.com/go-chi/chi"
)

func TestChartList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Get("/charts", list)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get charts with pagination", func(t *testing.T) {
		theme := &model.Theme{
			Name:           "Theme sample",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Theme{}).Create(&theme)

		chartOne := &model.Chart{
			Title:          "Sample-1",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}
		chartTwo := &model.Chart{
			Title:          "Sample-2",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
		}
		charts := []model.Chart{}
		total := 0

		config.DB.Model(&model.Chart{}).Create(&chartOne)
		config.DB.Model(&model.Chart{}).Create(&chartTwo)
		config.DB.Model(&model.Chart{}).Where(&model.Chart{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&charts)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "GET", "/charts?limit=1&page=2", nil, headers)

		respBody := (resp).(map[string]interface{})

		nodes := (respBody["nodes"]).([]interface{})
		chart := (nodes[0]).(map[string]interface{})
		gotTotal := (respBody["total"]).(float64)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if chart["title"] != charts[0].Title {
			t.Errorf("handler returned wrong title: got %v want %v", chart["title"], charts[0].Title)
		}

		if int(gotTotal) != total {
			t.Errorf("handler returned wrong total: got %v want %v",
				int(gotTotal), total)
		}

	})

}
