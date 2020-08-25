package category

import (
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"

	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete category by id
// @Summary Delete a category
// @Description Delete category by ID
// @Tags Category
// @ID delete-category-by-id
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param category_id path string true "Category ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /categories/{category_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

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

	result := &model.Category{}

	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Category{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// check if category is associated with charts
	category := new(model.Category)
	category.ID = uint(id)
	totAssociated := config.DB.Model(category).Association("Charts").Count()

	if totAssociated != 0 {
		errorx.Render(w, errorx.Parser(errorx.CannotSaveChanges()))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
