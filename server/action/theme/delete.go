package theme

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"

	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete theme by id
// @Summary Delete a theme
// @Description Delete theme by ID
// @Tags Theme
// @ID delete-theme-by-id
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param theme_id path string true "Theme ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /themes/{theme_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	themeID := chi.URLParam(r, "theme_id")
	id, err := strconv.Atoi(themeID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	result := &model.Theme{}

	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Theme{
		SpaceID: uint(sID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// check if theme is associated with charts
	uintID := uint(id)
	var totAssociated int64
	config.DB.Model(&model.Chart{}).Where(&model.Chart{
		ThemeID: &uintID,
	}).Count(&totAssociated)

	if totAssociated != 0 {
		loggerx.Error(errors.New("theme is associated with chart"))
		errorx.Render(w, errorx.Parser(errorx.CannotDelete("theme", "chart")))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
