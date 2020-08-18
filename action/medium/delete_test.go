package medium

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi"
	"gopkg.in/h2non/gock.v1"
)

func TestMediumDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/media", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid medium id", func(t *testing.T) {

		e.DELETE("/media/{medium_id}").
			WithPath("medium_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)

	})

	t.Run("medium record not found", func(t *testing.T) {
		e.DELETE("/media/{medium_id}").
			WithPath("medium_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("check medium associated with other entity", func(t *testing.T) {

		medium := model.Medium{
			Name:           "Entertainment",
			OrganisationID: 1,
		}

		theme := &model.Theme{
			Name:           "Light theme",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Medium{}).Create(&medium)
		config.DB.Model(&model.Theme{}).Create(&theme)

		t.Log(medium)

		chart := &model.Chart{
			Title:            "Bar chart",
			OrganisationID:   1,
			ThemeID:          theme.Base.ID,
			FeaturedMediumID: medium.Base.ID,
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		t.Log(chart)

		e.DELETE("/media/{medium_id}").
			WithPath("medium_id", medium.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("medium record deleted", func(t *testing.T) {
		medium := &model.Medium{
			Name:           "Cricket",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Medium{}).Create(&medium)

		e.DELETE("/media/{medium_id}").
			WithPath("medium_id", medium.Base.ID).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
