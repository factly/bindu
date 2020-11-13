package chart

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// update - Update chart by id
// @Summary Update a chart by id
// @Description Update chart by ID
// @Tags Chart
// @ID update-chart-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param chart_id path string true "Chart ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Chart body chart false "Chart"
// @Success 200 {object} model.Chart
// @Router /charts/{chart_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	chartID := chi.URLParam(r, "chart_id")
	id, err := strconv.Atoi(chartID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	chart := &chart{}
	err = json.NewDecoder(r.Body).Decode(&chart)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(chart)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Chart{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Chart{
		OrganisationID: uint(oID),
	}).Preload("Tags").Preload("Categories").First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// Fetching old and new tags related to chart
	oldTags := result.Tags
	newTags := make([]model.Tag, 0)
	config.DB.Model(&model.Tag{}).Where(chart.TagIDs).Find(&newTags)

	// Fetching old and new categories related to chart
	oldCategories := result.Categories
	newCategories := make([]model.Category, 0)
	config.DB.Model(&model.Category{}).Where(chart.CategoryIDs).Find(&newCategories)

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Chart{})
	tableName := stmt.Schema.Table

	var chartSlug string

	if result.Slug == chart.Slug {
		chartSlug = result.Slug
	} else if chart.Slug != "" && slug.Check(chart.Slug) {
		chartSlug = slug.Approve(chart.Slug, oID, tableName)
	} else {
		chartSlug = slug.Approve(slug.Make(chart.Title), oID, tableName)
	}

	// Deleting old associations
	if len(oldTags) > 0 {
		config.DB.Model(&result).Association("Tags").Delete(oldTags)
	}
	if len(oldCategories) > 0 {
		config.DB.Model(&result).Association("Categories").Delete(oldCategories)
	}

	if len(newTags) == 0 {
		newTags = nil
	}
	if len(newCategories) == 0 {
		newCategories = nil
	}

	config.DB.Model(&result).Updates(model.Chart{
		Title:            chart.Title,
		Slug:             chartSlug,
		DataURL:          chart.DataURL,
		Description:      chart.Description,
		Status:           chart.Status,
		FeaturedMediumID: chart.FeaturedMediumID,
		Config:           chart.Config,
		ThemeID:          chart.ThemeID,
		PublishedDate:    chart.PublishedDate,
		Tags:             newTags,
		Categories:       newCategories,
	}).Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
