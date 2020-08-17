package chart

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/paginationx"
	"github.com/factly/x/renderx"
)

// list response
type paging struct {
	Total int           `json:"total"`
	Nodes []model.Chart `json:"nodes"`
}

// list - Get all charts
// @Summary Show all charts
// @Description Get all charts
// @Tags Chart
// @ID get-all-charts
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {array} model.Chart
// @Router /charts [get]
func list(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	result := paging{}
	result.Nodes = make([]model.Chart, 0)

	offset, limit := paginationx.Parse(r.URL.Query())

	config.DB.Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").Model(&model.Chart{}).Where(&model.Chart{
		OrganisationID: uint(oID),
	}).Count(&result.Total).Order("id desc").Offset(offset).Limit(limit).Find(&result.Nodes)

	renderx.JSON(w, http.StatusOK, result)
}
