package category

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
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// update - Update category by id
// @Summary Update a category by id
// @Description Update category by ID
// @Tags Category
// @ID update-category-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param category_id path string true "Category ID"
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

	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	uID, err := middlewarex.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
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
		SpaceID: uint(sID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Category{})
	tableName := stmt.Schema.Table

	var categorySlug string

	if result.Slug == category.Slug {
		categorySlug = result.Slug
	} else if category.Slug != "" && slugx.Check(category.Slug) {
		categorySlug = slugx.Approve(&config.DB, category.Slug, sID, tableName)
	} else {
		categorySlug = slugx.Approve(&config.DB, slugx.Make(category.Name), sID, tableName)
	}

	// Check if category with same name exist
	if category.Name != result.Name && util.CheckCategoryName(uint(sID), category.Name, category.IsForTemplate) {
		loggerx.Error(errors.New(`category with same name exist`))
		errorx.Render(w, errorx.Parser(errorx.SameNameExist()))
		return
	}

	tx := config.DB.Begin()
	tx.Model(&result).Select("IsForTemplate").Updates(model.Category{IsForTemplate: category.IsForTemplate})
	tx.Model(&result).Updates(model.Category{
		Base:        config.Base{UpdatedByID: uint(uID)},
		Name:        category.Name,
		Slug:        categorySlug,
		Description: category.Description,
	}).First(&result)

	tx.Commit()

	renderx.JSON(w, http.StatusOK, result)
}
