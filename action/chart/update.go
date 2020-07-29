package chart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
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
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	chart := &chart{}
	json.NewDecoder(r.Body).Decode(&chart)

	result := &model.Chart{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Chart{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	var chartSlug string

	if result.Slug == chart.Slug {
		chartSlug = result.Slug
	} else if chart.Slug != "" && slug.Check(chart.Slug) {
		chartSlug = slug.Approve(chart.Slug, oID, config.DB.NewScope(&model.Chart{}).TableName())
	} else {
		chartSlug = slug.Approve(slug.Make(chart.Title), oID, config.DB.NewScope(&model.Chart{}).TableName())
	}

	// check themes, templates & medium belong to same organisation or not
	err = chart.CheckOrganisation(config.DB)
	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	config.DB.Model(&result).Updates(model.Chart{
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
	}).Preload("Medium").Preload("Template").Preload("Theme").First(&result).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
