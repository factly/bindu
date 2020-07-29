package template

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

// create - Create template
// @Summary Create template
// @Description Create template
// @Tags Template
// @ID add-template
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Template body template true "Template Object"
// @Success 201 {object} model.Template
// @Failure 400 {array} string
// @Router /templates [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	template := &template{}

	json.NewDecoder(r.Body).Decode(&template)

	validationError := validationx.Check(template)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var templateSlug string
	if template.Slug != "" && slug.Check(template.Slug) {
		templateSlug = template.Slug
	} else {
		templateSlug = slug.Make(template.Name)
	}

	result := &model.Template{
		Name:           template.Name,
		Slug:           slug.Approve(templateSlug, oID, config.DB.NewScope(&model.Template{}).TableName()),
		Description:    template.Description,
		URL:            template.URL,
		OrganisationID: uint(oID),
	}

	err = config.DB.Model(&model.Template{}).Create(&result).Error

	if err != nil {
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
