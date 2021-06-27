package organisation

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
)

type orgWithPermissions struct {
	model.Organisation
	Permission *model.OrganisationPermission `json:"permission"`
}

type paging struct {
	Nodes []orgWithPermissions `json:"nodes"`
	Total int64                `json:"total"`
}

// list - Get all organisation permissions
// @Summary Show all organisation permissions
// @Description Get all organisation permissions
// @Tags Organisation_Permissions
// @ID get-all-org-permissions
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param q query string false "Query"
// @Success 200 {array} paging
// @Router /permissions/organisations [get]
func list(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("q")

	result := paging{}
	result.Nodes = make([]orgWithPermissions, 0)

	permissionList := make([]model.OrganisationPermission, 0)

	config.DB.Model(&model.OrganisationPermission{}).Find(&permissionList)

	permissionMap := make(map[uint]model.OrganisationPermission)
	for _, permission := range permissionList {
		permissionMap[permission.OrganisationID] = permission
	}

	allOrgMap, err := util.GetAllOrganisationsMap(searchQuery)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	for oid, organisation := range allOrgMap {
		owp := orgWithPermissions{}
		owp.Organisation = organisation
		if per, found := permissionMap[oid]; found {
			owp.Permission = &per
		}

		result.Nodes = append(result.Nodes, owp)
	}

	result.Total = int64(len(result.Nodes))

	renderx.JSON(w, http.StatusOK, result)
}
