package theme

import (
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// details - Get theme by id
// @Summary Show a theme by id
// @Description Get theme by ID
// @Tags Theme
// @ID get-theme-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param theme_id path string true "Theme ID"
// @Success 200 {object} model.Theme
// @Router /themes/{theme_id} [get]
func details(w http.ResponseWriter, r *http.Request) {

	themeID := chi.URLParam(r, "theme_id")
	id, err := strconv.Atoi(themeID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	result := &model.Theme{}

	result.ID = uint(id)

	err = config.DB.Model(&model.Theme{}).Where(&model.Theme{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}
