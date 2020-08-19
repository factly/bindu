package tag

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

// create - Create tag
// @Summary Create tag
// @Description Create tag
// @Tags Tag
// @ID add-tag
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Tag body tag true "Tag Object"
// @Success 201 {object} model.Tag
// @Failure 400 {array} string
// @Router /tags [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	tag := &tag{}

	json.NewDecoder(r.Body).Decode(&tag)

	validationError := validationx.Check(tag)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var tagSlug string
	if tag.Slug != "" && slug.Check(tag.Slug) {
		tagSlug = tag.Slug
	} else {
		tagSlug = slug.Make(tag.Name)
	}

	result := &model.Tag{
		Name:           tag.Name,
		Slug:           slug.Approve(tagSlug, oID, config.DB.NewScope(&model.Tag{}).TableName()),
		Description:    tag.Description,
		OrganisationID: uint(oID),
	}

	err = config.DB.Model(&model.Tag{}).Create(&result).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
