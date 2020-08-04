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

// delete - Delete theme by id
// @Summary Delete a theme
// @Description Delete theme by ID
// @Tags Theme
// @ID delete-theme-by-id
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param theme_id path string true "Theme ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /themes/{theme_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

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

	// check record exists or not
	err = config.DB.Where(&model.Theme{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// check if theme is associated with charts
	var totAssociated int
	config.DB.Model(&model.Chart{}).Where(&model.Chart{
		ThemeID: uint(id),
	}).Count(&totAssociated)

	if totAssociated != 0 {
		errorx.Render(w, errorx.Parser(errorx.CannotSaveChanges()))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
