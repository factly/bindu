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

func TestMediumDelete(t *testing.T) {
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
		_, statusCode := test.Request(t, ts, "DELETE", "/media/invalid_id", nil, headers)

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
		_, statusCode := test.Request(t, ts, "DELETE", "/media/100", nil, headers)

		if statusCode != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusNotFound)
		}
	})

	t.Run("check medium associated with other entity", func(t *testing.T) {

		medium := &model.Medium{
			Name:           "sample",
			OrganisationID: 1,
		}

		theme := &model.Theme{
			Name: "Sample Theme",
		}

		config.DB.Model(&model.Medium{}).Create(&medium)
		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:            "Sample chart",
			OrganisationID:   1,
			FeaturedMediumID: medium.Base.ID,
			ThemeID:          theme.Base.ID,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/media/", medium.Base.ID), nil, headers)

		if statusCode != http.StatusUnprocessableEntity {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusUnprocessableEntity)
		}
	})

	t.Run("medium record deleted", func(t *testing.T) {
		headers := map[string]string{
			"X-Organisation": "1",
			"X-User":         "1",
		}
		result := &model.Medium{
			Name:           "testing",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Medium{}).Create(&result)

		_, statusCode := test.Request(t, ts, "DELETE", fmt.Sprint("/media/", result.Base.ID), nil, headers)

		if statusCode != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				statusCode, http.StatusOK)
		}

	})

}
