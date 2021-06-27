package space

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/paginationx"
	"github.com/factly/x/renderx"
)

type paging struct {
	Nodes []model.SpacePermissionRequest `json:"nodes"`
	Total int64                          `json:"total"`
}

// list - Get all space permissions requests
// @Summary Show all space permissions requests
// @Description Get all space permissions requests
// @Tags Space_Permissions_Request
// @ID get-all-space-permissions-requests
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param status query string false "Status"
// @Success 200 {array} paging
// @Router /requests/spaces [get]
func list(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "pending"
	}

	offset, limit := paginationx.Parse(r.URL.Query())

	result := paging{}
	result.Nodes = make([]model.SpacePermissionRequest, 0)

	tx := config.DB.Model(&model.SpacePermissionRequest{})

	if status != "all" {
		tx.Where("status = ?", status)
	}

	tx.Count(&result.Total).Offset(offset).Limit(limit).Find(&result.Nodes)

	renderx.JSON(w, http.StatusOK, result)
}
