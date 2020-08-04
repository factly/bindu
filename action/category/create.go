package category

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

// create - Create category
// @Summary Create category
// @Description Create category
// @Tags Category
// @ID add-category
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Category body category true "Category Object"
// @Success 201 {object} model.Category
// @Failure 400 {array} string
// @Router /categories [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	category := &category{}

	json.NewDecoder(r.Body).Decode(&category)

	validationError := validationx.Check(category)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var categorySlug string
	if category.Slug != "" && slug.Check(category.Slug) {
		categorySlug = category.Slug
	} else {
		categorySlug = slug.Make(category.Name)
	}

	result := &model.Category{
		Name:           category.Name,
		Slug:           slug.Approve(categorySlug, oID, config.DB.NewScope(&model.Category{}).TableName()),
		Description:    category.Description,
		OrganisationID: uint(oID),
	}

	err = config.DB.Model(&model.Category{}).Create(&result).Error

	if err != nil {
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
