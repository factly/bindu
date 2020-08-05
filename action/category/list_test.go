package category

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

func TestCategoryList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Get("/categories", list)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get categories with pagination", func(t *testing.T) {
		categoryOne := &model.Category{
			Name:           "Sample-1",
			OrganisationID: 1,
		}
		categoryTwo := &model.Category{
			Name:           "Sample-2",
			OrganisationID: 1,
		}
		categories := []model.Category{}
		total := 0

		config.DB.Model(&model.Category{}).Create(&categoryOne)
		config.DB.Model(&model.Category{}).Create(&categoryTwo)
		config.DB.Model(&model.Category{}).Where(&model.Category{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&categories)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "GET", "/categories?limit=1&page=2", nil, headers)

		respBody := (resp).(map[string]interface{})

		nodes := (respBody["nodes"]).([]interface{})
		category := (nodes[0]).(map[string]interface{})
		gotTotal := (respBody["total"]).(float64)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if category["name"] != categories[0].Name {
			t.Errorf("handler returned wrong title: got %v want %v", category["name"], categories[0].Name)
		}

		if int(gotTotal) != total {
			t.Errorf("handler returned wrong total: got %v want %v",
				int(gotTotal), total)
		}

	})

}
