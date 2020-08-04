package chart

import (
	"encoding/json"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
)

// create - Create chart
// @Summary Create chart
// @Description Create chart
// @Tags Chart
// @ID add-chart
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Chart body chart true "Chart Object"
// @Success 201 {object} model.Chart
// @Failure 400 {array} string
// @Router /charts [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	chart := &chart{}

	json.NewDecoder(r.Body).Decode(&chart)

	validationError := validationx.Check(chart)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var chartSlug string
	if chart.Slug != "" && slug.Check(chart.Slug) {
		chartSlug = chart.Slug
	} else {
		chartSlug = slug.Make(chart.Title)
	}

	result := &model.Chart{
		Title:            chart.Title,
		Slug:             slug.Approve(chartSlug, oID, config.DB.NewScope(&model.Chart{}).TableName()),
		DataURL:          chart.DataURL,
		Config:           chart.Config,
		Description:      chart.Description,
		Status:           chart.Status,
		FeaturedMediumID: chart.FeaturedMediumID,
		ThemeID:          chart.ThemeID,
		PublishedDate:    chart.PublishedDate,
		OrganisationID:   uint(oID),
	}

	// check themes & medium belong to same organisation or not
	err = chart.CheckOrganisation(config.DB)
	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	config.DB.Model(&model.Tag{}).Where(chart.TagIDs).Find(&result.Tags)
	config.DB.Model(&model.Category{}).Where(chart.CategoryIDs).Find(&result.Categories)

	err = config.DB.Model(&model.Chart{}).Set("gorm:association_autoupdate", false).Create(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	config.DB.Model(&model.Chart{}).Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").First(&result)

	renderx.JSON(w, http.StatusCreated, result)
}
