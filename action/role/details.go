package role

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factly/bindu-server/action/user"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// details - Get Role by ID
// @Summary Get Role by ID
// @Description Get Role by ID
// @Tags Role
// @ID get-role-by-id
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param role_id path string true "Role ID"
// @Success 200 {object} model.Role
// @Router /roles/{role_id} [get]
func details(w http.ResponseWriter, r *http.Request) {
	sID, err := util.GetSpace(r.Context())
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

	roleID := chi.URLParam(r, "role_id")
	if roleID == "" {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	ketoRoleID := fmt.Sprint("roles:org:", oID, ":app:bindu:space:", sID, ":", roleID)

	resp, err := util.Request("GET", viper.GetString("keto_url")+"/engines/acp/ory/regex/roles/"+ketoRoleID, nil)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	if resp.StatusCode == http.StatusNotFound {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	var ketoRole model.KetoRole

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ketoRole)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	/* User req */
	userMap := user.Mapper(oID, uID)

	result := Mapper(ketoRole, userMap)

	renderx.JSON(w, http.StatusOK, result)

}
