package theme

import (
	"encoding/json"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
)

// create - Create theme
// @Summary Create theme
// @Description Create theme
// @Tags Theme
// @ID add-theme
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Theme body theme true "Theme Object"
// @Success 201 {object} model.Theme
// @Failure 400 {array} string
// @Router /themes [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	theme := &theme{}

	json.NewDecoder(r.Body).Decode(&theme)

	validationError := validationx.Check(theme)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var themeSlug string
	if theme.Slug != "" && slug.Check(theme.Slug) {
		themeSlug = theme.Slug
	} else {
		themeSlug = slug.Make(theme.Name)
	}

	result := &model.Theme{
		Name:           theme.Name,
		Slug:           slug.Approve(themeSlug, oID, config.DB.NewScope(&model.Theme{}).TableName()),
		Description:    theme.Description,
		URL:            theme.URL,
		OrganisationID: uint(oID),
	}

	err = config.DB.Model(&model.Theme{}).Create(&result).Error

	if err != nil {
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
