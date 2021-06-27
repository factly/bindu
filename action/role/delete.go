package role

import (
	"fmt"
	"net/http"

	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// delete - Delete Role by ID
// @Summary Delete Role by ID
// @Description Delete Role by ID
// @Tags Role
// @ID delete-role-by-id
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param role_id path string true "Role ID"
// @Success 200
// @Router /roles/{role_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {
	sID, err := middlewarex.GetSpace(r.Context())
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

	roleID := chi.URLParam(r, "role_id")
	if roleID == "" {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	ketoRoleID := fmt.Sprint("roles:org:", oID, ":app:bindu:space:", sID, ":", roleID)

	resp, err := util.Request("DELETE", viper.GetString("keto_url")+"/engines/acp/ory/regex/roles/"+ketoRoleID, nil)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	if resp.StatusCode != http.StatusNoContent {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	renderx.JSON(w, http.StatusOK, nil)
}
