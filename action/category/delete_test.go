package category

import (
	"fmt"
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

func TestCategoryDelete(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/categories", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	t.Run("invalid category id", func(t *testing.T) {

		e.DELETE("/categories/invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)

	})

	t.Run("category record not found", func(t *testing.T) {
		e.DELETE("/categories/100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("check category associated with other entity", func(t *testing.T) {

		category := model.Category{
			Name:           "sample",
			OrganisationID: 1,
		}

		theme := &model.Theme{
			Name: "Sample Theme",
		}

		config.DB.Model(&model.Category{}).Create(&category)
		config.DB.Model(&model.Theme{}).Create(&theme)

		chart := &model.Chart{
			Title:          "Sample chart",
			OrganisationID: 1,
			ThemeID:        theme.Base.ID,
			Categories:     []model.Category{category},
		}

		config.DB.Model(&model.Chart{}).Create(&chart)

		e.DELETE(fmt.Sprint("/categories/", category.Base.ID)).
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)
	})

	t.Run("category record deleted", func(t *testing.T) {
		category := &model.Category{
			Name:           "Cricket",
			OrganisationID: 1,
		}

		config.DB.Model(&model.Category{}).Create(&category)

		e.DELETE(fmt.Sprint("/categories/", category.Base.ID)).
			WithHeaders(headers).
			Expect().
			Status(http.StatusOK)
	})

}
