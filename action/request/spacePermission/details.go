package spacePermission

import (
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// details - Get space permissions requests detail
// @Summary Show a space permissions requests detail
// @Description Get space permissions requests detail
// @Tags Space_Permissions_Request
// @ID get-space-permission-request-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param request_id path string true "Request ID"
// @Success 200 {object} model.SpacePermissionRequest
// @Router /requests/space-permissions/{request_id} [get]
func details(w http.ResponseWriter, r *http.Request) {
	requestID := chi.URLParam(r, "request_id")
	id, err := strconv.Atoi(requestID)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	result := model.SpacePermissionRequest{}
	result.ID = uint(id)

	err = config.DB.First(&result).Error
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}
