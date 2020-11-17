package organisation

import (
	"net/http"

	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
)

// list response
type paging struct {
	Total int64                `json:"total"`
	Nodes []model.Organisation `json:"nodes"`
}

// list - Get all organisations
// @Summary Show all organisations
// @Description Get all organisations
// @Tags Organisation
// @ID get-all-organisations
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {array} model.Organisation
// @Router /organisations [get]
func list(w http.ResponseWriter, r *http.Request) {

	orgs, err := util.RequestOrganisation(r)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	result := paging{}
	result.Nodes = make([]model.Organisation, 0)

	result.Nodes = orgs

	result.Total = int64(len(orgs))

	renderx.JSON(w, http.StatusOK, result)
}
