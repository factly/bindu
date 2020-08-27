package tag

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/paginationx"
	"github.com/factly/x/renderx"
)

// list response
type paging struct {
	Total int         `json:"total"`
	Nodes []model.Tag `json:"nodes"`
}

// list - Get all tags
// @Summary Show all tags
// @Description Get all tags
// @Tags Tag
// @ID get-all-tags
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {array} paging
// @Router /tags [get]
func list(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	result := paging{}
	result.Nodes = make([]model.Tag, 0)

	offset, limit := paginationx.Parse(r.URL.Query())

	config.DB.Model(&model.Tag{}).Where(&model.Tag{
		OrganisationID: uint(oID),
	}).Count(&result.Total).Order("id desc").Offset(offset).Limit(limit).Find(&result.Nodes)

	renderx.JSON(w, http.StatusOK, result)
}
