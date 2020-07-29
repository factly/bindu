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
		Subtitle:         chart.Subtitle,
		Slug:             slug.Approve(chartSlug, oID, config.DB.NewScope(&model.Chart{}).TableName()),
		URL:              chart.URL,
		Description:      chart.Description,
		Status:           chart.Status,
		FeaturedMediumID: chart.FeaturedMediumID,
		TemplateID:       chart.TemplateID,
		ThemeID:          chart.ThemeID,
		PublishedDate:    chart.PublishedDate,
		OrganisationID:   uint(oID),
	}

	// check themes, templates & medium belong to same organisation or not
	err = chart.CheckOrganisation(config.DB)
	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	err = config.DB.Model(&model.Chart{}).Create(&result).Error

	if err != nil {
		return
	}

	config.DB.Model(&model.Chart{}).Preload("Medium").Preload("Template").Preload("Theme").First(&result)

	renderx.JSON(w, http.StatusCreated, result)
}
