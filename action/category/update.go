package category

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
)

// update - Update category by id
// @Summary Update a category by id
// @Description Update category by ID
// @Tags Category
// @ID update-category-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param category_id path string true "Category ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Category body category false "Category"
// @Success 200 {object} model.Category
// @Router /categories/{category_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "category_id")
	id, err := strconv.Atoi(categoryID)

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
	category := &category{}
	err = json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(category)

	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Category{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Category{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	var categorySlug string

	if result.Slug == category.Slug {
		categorySlug = result.Slug
	} else if category.Slug != "" && slug.Check(category.Slug) {
		categorySlug = slug.Approve(category.Slug, oID, config.DB.NewScope(&model.Category{}).TableName())
	} else {
		categorySlug = slug.Approve(slug.Make(category.Name), oID, config.DB.NewScope(&model.Category{}).TableName())
	}

	config.DB.Model(&result).Updates(model.Category{
		Name:        category.Name,
		Slug:        categorySlug,
		Description: category.Description,
	}).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
