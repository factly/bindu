package tag

import (
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

func TestTagList(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/tags", Router())

	ts := httptest.NewServer(r)
	gock.New(ts.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer ts.Close()

	t.Run("get tags with pagination", func(t *testing.T) {
		tagOne := &model.Tag{
			Name:           "Tag-1",
			OrganisationID: 1,
		}
		tagTwo := &model.Tag{
			Name:           "Tag-2",
			OrganisationID: 1,
		}
		tags := []model.Tag{}
		total := 0

		config.DB.Model(&model.Tag{}).Create(&tagOne)
		config.DB.Model(&model.Tag{}).Create(&tagTwo)
		config.DB.Model(&model.Tag{}).Where(&model.Tag{
			OrganisationID: 1,
		}).Count(&total).Order("id desc").Offset(1).Limit(1).Find(&tags)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		resp, statusCode := test.Request(t, ts, "GET", "/tags?limit=1&page=2", nil, headers)

		respBody := (resp).(map[string]interface{})

		nodes := (respBody["nodes"]).([]interface{})
		tag := (nodes[0]).(map[string]interface{})
		gotTotal := (respBody["total"]).(float64)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

		if tag["name"] != tags[0].Name {
			t.Errorf("handler returned wrong title: got %v want %v", tag["name"], tags[0].Name)
		}

		if int(gotTotal) != total {
			t.Errorf("handler returned wrong total: got %v want %v",
				int(gotTotal), total)
		}

	})

}
