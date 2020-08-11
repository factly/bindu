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

func TestTagDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Delete("/tags/{tag_id}", delete)

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("invalid tag id", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", "/tags/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "DELETE", "/tags/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("check tag associated with other entity", func(t *testing.T) {

		tag := model.Tag{
			Name:           "sample",
			OrganisationID: 1,
		}

		theme := &model.Theme{
			Name: "Sample Theme",
		}

		config.DB.Model(&model.Tag{}).Create(&tag)
		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:          "Sample chart",
			OrganisationID: 1,
			ThemeID:        theme.Base.ID,
			Tags:           []model.Tag{tag},
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/tags/", tag.Base.ID), nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("tag record deleted", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		result := &model.Tag{
			Name:           "Football",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Tag{}).Create(&result)

		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/tags/", result.Base.ID), nil, headers)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

}
