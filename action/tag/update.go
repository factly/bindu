package tag

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// update - Update tag by id
// @Summary Update a tag by id
// @Description Update tag by ID
// @Tags Tag
// @ID update-tag-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param tag_id path string true "Tag ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Tag body tag false "Tag"
// @Success 200 {object} model.Tag
// @Router /tags/{tag_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
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
	tag := &tag{}
	json.NewDecoder(r.Body).Decode(&tag)

	result := &model.Tag{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Tag{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	var tagSlug string

	if result.Slug == tag.Slug {
		tagSlug = result.Slug
	} else if tag.Slug != "" && slug.Check(tag.Slug) {
		tagSlug = slug.Approve(tag.Slug, oID, config.DB.NewScope(&model.Tag{}).TableName())
	} else {
		tagSlug = slug.Approve(slug.Make(tag.Name), oID, config.DB.NewScope(&model.Tag{}).TableName())
	}

	config.DB.Model(&result).Updates(model.Tag{
		Name:        tag.Name,
		Slug:        tagSlug,
		Description: tag.Description,
	}).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
