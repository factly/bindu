package theme

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

func TestThemeList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Get("/themes", list)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get themes with pagination", func(t *testing.T) {
		themeOne := &model.Theme{
			Name:           "Sample-1",
			OrganisationID: 1,
		}
		themeTwo := &model.Theme{
			Name:           "Sample-2",
			OrganisationID: 1,
		}
		themes := []model.Theme{}
		total := 0

		config.DB.Model(&model.Theme{}).Create(&themeOne)
		config.DB.Model(&model.Theme{}).Create(&themeTwo)
		config.DB.Model(&model.Theme{}).Where(&model.Theme{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&themes)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "GET", "/themes?limit=1&page=2", nil, headers)

		respBody := (resp).(map[string]interface{})

		nodes := (respBody["nodes"]).([]interface{})
		theme := (nodes[0]).(map[string]interface{})
		gotTotal := (respBody["total"]).(float64)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if theme["name"] != themes[0].Name {
			t.Errorf("handler returned wrong title: got %v want %v", theme["name"], themes[0].Name)
		}

		if int(gotTotal) != total {
			t.Errorf("handler returned wrong status code: got %v want %v",
				int(gotTotal), total)
		}

	})

}
