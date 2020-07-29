package chart

import (
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"

	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete chart by id
// @Summary Delete a chart
// @Description Delete chart by ID
// @Tags Chart
// @ID delete-chart-by-id
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param chart_id path string true "Chart ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /charts/{chart_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

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

	result := &model.Chart{}

	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Chart{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
