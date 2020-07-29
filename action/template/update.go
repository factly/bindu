package template

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

// update - Update template by id
// @Summary Update a template by id
// @Description Update template by ID
// @Tags Template
// @ID update-template-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param template_id path string true "Template ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Template body template false "Template"
// @Success 200 {object} model.Template
// @Router /templates/{template_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	templateID := chi.URLParam(r, "template_id")
	id, err := strconv.Atoi(templateID)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	template := &template{}
	json.NewDecoder(r.Body).Decode(&template)

	result := &model.Template{}
	result.ID = uint(id)

	// check record exists or not
	err = config.DB.Where(&model.Template{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	var templateSlug string

	if result.Slug == template.Slug {
		templateSlug = result.Slug
	} else if template.Slug != "" && slug.Check(template.Slug) {
		templateSlug = slug.Approve(template.Slug, oID, config.DB.NewScope(&model.Template{}).TableName())
	} else {
		templateSlug = slug.Approve(slug.Make(template.Name), oID, config.DB.NewScope(&model.Template{}).TableName())
	}

	config.DB.Model(&result).Updates(model.Template{
		Name:        template.Name,
		Slug:        templateSlug,
		URL:         template.URL,
		Description: template.Description,
	}).First(&result)

	renderx.JSON(w, http.StatusOK, result)
}
