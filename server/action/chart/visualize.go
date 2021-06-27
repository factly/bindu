package chart

import (
	"encoding/json"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/go-chi/chi"
)

// Visualize - Get chart by id
// @Summary Show a chart by id
// @Description Get chart by ID
// @Tags Chart
// @ID get-chart-visualization-by-id
// @Produce json
// @Param chart_id path string true "Chart ID"
// @Success 200
// @Router /charts/visualization/{chart_id} [get]
func Visualize(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "chart_id")
	if id == "" {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := &model.Chart{}
	result.ID = id

	err := config.DB.Model(&model.Chart{}).Where(&model.Chart{
		IsPublic: true,
	}).Where("published_date IS NOT NULL").Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	specMap := util.Unmarshal(result.Config)
	jsonBytes, _ := json.Marshal(specMap)

	jsonSpecString := string(jsonBytes)

	err = util.Template.ExecuteTemplate(w, "chart.gohtml", map[string]interface{}{
		"chart": result,
		"spec":  jsonSpecString,
	})
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
}
