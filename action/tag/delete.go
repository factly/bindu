package tag

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

// delete - Delete tag by id
// @Summary Delete a tag
// @Description Delete tag by ID
// @Tags Tag
// @ID delete-tag-by-id
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param tag_id path string true "Tag ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /tags/{tag_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	tagID := chi.URLParam(r, "tag_id")
	id, err := strconv.Atoi(tagID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	result := &model.Tag{}

	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Tag{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// check if tag is associated with charts
	tag := new(model.Tag)
	tag.ID = uint(id)
	totAssociated := config.DB.Model(tag).Association("Charts").Count()

	if totAssociated != 0 {
		errorx.Render(w, errorx.Parser(errorx.CannotSaveChanges()))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
