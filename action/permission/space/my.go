package space

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
)

// my - Get my Space permissions
// @Summary Show a my Space permissions
// @Description Get my Space permissions
// @Tags Space_Permissions
// @ID get-my-space-permissions
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Success 200 {object} model.SpacePermission
// @Router /permissions/spaces/my [get]
func my(w http.ResponseWriter, r *http.Request) {
	sID, err := util.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	result := model.SpacePermission{}

	err = config.DB.Model(&model.SpacePermission{}).Where(&model.SpacePermission{
		SpaceID: uint(sID),
	}).Preload("Space").First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}
