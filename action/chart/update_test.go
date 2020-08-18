package chart

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

func TestChartUpdate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/charts", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	// create httpexpect instance
	e := httpexpect.New(t, testServer.URL)

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	category := &model.Category{
		Name:           "Sports",
		OrganisationID: 1,
	}

	tag := &model.Tag{
		Name:           "Agriculture",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Tag{}).Create(&tag)
	config.DB.Model(&model.Category{}).Create(&category)
	config.DB.Model(&model.Theme{}).Create(&theme)

	result := &model.Chart{
		Title:          "Test",
		Slug:           "maps",
		ThemeID:        theme.Base.ID,
		OrganisationID: 1,
	}

	config.DB.Model(&model.Tag{}).Where(tag.Base.ID).Find(&result.Tags)
	config.DB.Model(&model.Category{}).Where(category.Base.ID).Find(&result.Categories)

	config.DB.Model(&model.Chart{}).Set("gorm:association_autoupdate", false).Create(&result)

	t.Run("invalid chart id", func(t *testing.T) {
		e.PUT("/charts/{chart_id}").
			WithPath("chart_id", "invalid_id").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("chart record not found", func(t *testing.T) {
		e.PUT("/charts/{chart_id}").
			WithPath("chart_id", "100").
			WithHeaders(headers).
			Expect().
			Status(http.StatusNotFound)
	})

	t.Run("update chart", func(t *testing.T) {

		body := &chart{
			Title:          "Maps",
			Slug:           "maps",
			ThemeID:        theme.Base.ID,
			OrganisationID: 1,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}

		resObj := e.PUT("/charts/{chart_id}").
			WithPath("chart_id", result.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("title").String().Equal("Maps")

	})

	t.Run("update chart by id with empty slug", func(t *testing.T) {

		body := &model.Chart{
			Title: "Pie",
			Slug:  "",
		}

		resObj := e.PUT("/charts/{chart_id}").
			WithPath("chart_id", result.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("title").String().Equal("Pie")
		resObj.Value("slug").String().Equal("pie")

	})

	t.Run("update chart with different slug", func(t *testing.T) {

		body := &model.Chart{
			Title: "Chart test",
			Slug:  "map-sample",
		}

		resObj := e.PUT("/charts/{chart_id}").
			WithPath("chart_id", result.Base.ID).
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusOK).JSON().Object()

		resObj.Value("title").String().Equal("Chart test")
		resObj.Value("slug").String().Equal("map-sample")

	})

}
