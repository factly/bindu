package chart

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// Spec - Get chart spec by id
// @Summary Show a spec chart by id
// @Description Get spec chart by ID
// @Tags Chart
// @ID get-chart-spec-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param chart_id path string true "Chart ID"
// @Success 200 {object} model.Chart
// @Router /charts/{chart_id} [get]
func Spec(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "chart_id")
	if id == "" {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &model.Chart{}
	result.ID = id

	err := config.DB.Model(&model.Chart{}).Where(&model.Chart{
		IsPublic: true,
	}).Where("published_date IS NOT NULL").First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	specMap := util.Unmarshal(result.Config)

	renderx.JSON(w, http.StatusOK, specMap)
}
