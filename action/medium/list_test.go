package medium

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

func TestMediumList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.GenerateOrganisation).Get("/media", list)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("get media with pagination", func(t *testing.T) {
		mediumOne := &model.Medium{
			Name:           "Sample-1",
			OrganisationID: 1,
		}
		mediumTwo := &model.Medium{
			Name:           "Sample-2",
			OrganisationID: 1,
		}
		media := []model.Medium{}
		total := 0

		config.DB.Model(&model.Medium{}).Create(&mediumOne)
		config.DB.Model(&model.Medium{}).Create(&mediumTwo)
		config.DB.Model(&model.Medium{}).Where(&model.Medium{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&media)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "GET", "/media?limit=1&page=2", nil, headers)

		respBody := (resp).(map[string]interface{})

		nodes := (respBody["nodes"]).([]interface{})
		medium := (nodes[0]).(map[string]interface{})
		gotTotal := (respBody["total"]).(float64)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if medium["name"] != media[0].Name {
			t.Errorf("handler returned wrong title: got %v want %v", medium["name"], media[0].Name)
		}

		if int(gotTotal) != total {
			t.Errorf("handler returned wrong status code: got %v want %v",
				int(gotTotal), total)
		}

	})

}
