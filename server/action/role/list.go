package role

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/factly/bindu-server/action/user"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/paginationx"
	"github.com/factly/x/renderx"
	"github.com/spf13/viper"
)

type paging struct {
	Total int          `json:"total"`
	Nodes []model.Role `json:"nodes"`
}

// list - Get all roles
// @Summary Get all roles
// @Description Get all roles
// @Tags Role
// @ID get-all-roles
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {object} paging
// @Router /roles [get]
func list(w http.ResponseWriter, r *http.Request) {
	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	uID, err := middlewarex.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	resp, err := util.Request("GET", viper.GetString("keto_url")+"/engines/acp/ory/regex/roles", nil)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	var ketoRoles []model.KetoRole

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ketoRoles)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	prefixID := fmt.Sprint("roles:org:", oID, ":app:bindu:space:", sID, ":")

	onlySpaceRoles := make([]model.KetoRole, 0)

	for _, role := range ketoRoles {
		if strings.HasPrefix(role.ID, prefixID) {
			onlySpaceRoles = append(onlySpaceRoles, role)
		}
	}

	offset, limit := paginationx.Parse(r.URL.Query())

	total := len(onlySpaceRoles)
	lowerLimit := offset
	upperLimit := offset + limit
	if offset > total {
		lowerLimit = 0
		upperLimit = 0
	} else if offset+limit > total {
		lowerLimit = offset
		upperLimit = total
	}

	onlySpaceRoles = onlySpaceRoles[lowerLimit:upperLimit]

	/* User req */
	userMap := user.Mapper(oID, uID)

	pageRoles := make([]model.Role, 0)

	for _, each := range onlySpaceRoles {
		pageRoles = append(pageRoles, Mapper(each, userMap))
	}

	var result paging
	result.Nodes = pageRoles
	result.Total = total

	renderx.JSON(w, http.StatusOK, result)
}
