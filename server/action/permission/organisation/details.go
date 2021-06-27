package organisation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/spf13/viper"
)

type orgPermissionRes struct {
	model.OrganisationPermission
	SpacePermissions []model.SpacePermission `json:"space_permissions"`
	IsAdmin          bool                    `json:"is_admin,omitempty"`
}

// details - Get my organisation permissions
// @Summary Show a my organisation permissions
// @Description Get my organisation permissions
// @Tags Organisation_Permissions
// @ID get-org-permission-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Success 200 {object} orgPermissionRes
// @Router /permissions/organisations/my [get]
func details(w http.ResponseWriter, r *http.Request) {
	oID, err := util.GetOrganisation(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	result := orgPermissionRes{}
	err = config.DB.Model(&model.OrganisationPermission{}).Where(&model.OrganisationPermission{
		OrganisationID: uint(oID),
	}).First(&result.OrganisationPermission).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	resp, _ := util.Request("GET", viper.GetString("keto_url")+"/engines/acp/ory/regex/policies/app:bindu:superorg", nil)

	if resp.StatusCode == http.StatusOK {
		var policy model.KetoPolicy
		err = json.NewDecoder(resp.Body).Decode(&policy)
		if err == nil && len(policy.Subjects) > 0 && policy.Subjects[0] == fmt.Sprint(oID) {
			result.IsAdmin = true
		}
	}

	// Get all spaces of organisation
	spaceList := make([]model.Space, 0)
	config.DB.Model(&model.Space{}).Where(&model.Space{
		OrganisationID: oID,
	}).Find(&spaceList)

	spaceIDs := make([]uint, 0)
	for _, space := range spaceList {
		spaceIDs = append(spaceIDs, space.ID)
	}

	// Fetch all the spaces's permissions
	err = config.DB.Model(&model.SpacePermission{}).Where("space_id IN (?)", spaceIDs).Find(&result.SpacePermissions).Error
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}
