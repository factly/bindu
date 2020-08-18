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

func TestChartCreate(t *testing.T) {
	r := chi.NewRouter()

	r.With(util.CheckUser, util.CheckOrganisation).Mount("/charts", Router())

	testServer := httptest.NewServer(r)
	gock.New(testServer.URL).EnableNetworking().Persist()
	defer gock.DisableNetworking()
	defer testServer.Close()

	e := httpexpect.New(t, testServer.URL)

	category := &model.Category{
		Name:           "Sports",
		OrganisationID: 1,
	}

	tag := &model.Tag{
		Name:           "Agriculture",
		OrganisationID: 1,
	}

	theme := &model.Theme{
		Name:           "Theme sample",
		OrganisationID: 1,
	}

	config.DB.Model(&model.Theme{}).Create(&theme)
	config.DB.Model(&model.Tag{}).Create(&tag)
	config.DB.Model(&model.Category{}).Create(&category)

	t.Run("Unprocessable chart", func(t *testing.T) {
		e.POST("/charts").
			WithHeaders(headers).
			Expect().
			Status(http.StatusUnprocessableEntity)

	})

	t.Run("create chart", func(t *testing.T) {
		body := chart{
			Title:   "Pie Chart",
			Slug:    "pie-chart",
			ThemeID: theme.Base.ID,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}
		resObj := e.POST("/charts").
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("title").String().Equal("Pie Chart")
		resObj.Value("slug").String().Equal("pie-chart")

	})

	t.Run("Invalid theme id, medium id, category ids & tag ids", func(t *testing.T) {

		body := chart{
			Title:            "Bar",
			ThemeID:          100,
			FeaturedMediumID: 100,
			TagIDs: []uint{
				100,
			},
			CategoryIDs: []uint{
				100,
			},
		}

		e.POST("/charts").
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusInternalServerError)

	})

	t.Run("create chart with slug is empty", func(t *testing.T) {

		body := chart{
			Title:   "Bar",
			Slug:    "bar",
			ThemeID: theme.Base.ID,
			TagIDs: []uint{
				tag.Base.ID,
			},
			CategoryIDs: []uint{
				category.Base.ID,
			},
		}

		resObj := e.POST("/charts").
			WithHeaders(headers).
			WithJSON(body).
			Expect().
			Status(http.StatusCreated).JSON().Object()

		resObj.Value("title").String().Equal("Bar")
		resObj.Value("slug").String().Equal("bar")

	})

}
