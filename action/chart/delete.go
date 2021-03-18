package chart

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"

	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete chart by id
// @Summary Delete a chart
// @Description Delete chart by ID
// @Tags Chart
// @ID delete-chart-by-id
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param chart_id path string true "Chart ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /charts/{chart_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "chart_id")
	if id == "" {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	result := &model.Chart{}
	result.ID = id

	// check record exists or not
	err = config.DB.Where(&model.Chart{
		SpaceID: uint(sID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
