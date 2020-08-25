package theme

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
)

// update - Update theme by id
// @Summary Update a theme by id
// @Description Update theme by ID
// @Tags Theme
// @ID update-theme-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param theme_id path string true "Theme ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Theme body theme false "Theme"
// @Success 200 {object} model.Theme
// @Router /themes/{theme_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	themeID := chi.URLParam(r, "theme_id")
	id, err := strconv.Atoi(themeID)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	theme := &theme{}

	err = json.NewDecoder(r.Body).Decode(&theme)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(theme)

	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Theme{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Theme{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	config.DB.Model(&result).Updates(model.Theme{
		Name:   theme.Name,
		Config: theme.Config,
	}).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
