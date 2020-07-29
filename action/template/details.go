package template

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

// details - Get template by id
// @Summary Show a template by id
// @Description Get template by ID
// @Tags Template
// @ID get-template-by-id
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param template_id path string true "Template ID"
// @Success 200 {object} model.Template
// @Router /templates/{template_id} [get]
func details(w http.ResponseWriter, r *http.Request) {

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

	result := &model.Template{}

	result.ID = uint(id)

	err = config.DB.Model(&model.Template{}).Where(&model.Template{
		OrganisationID: uint(oID),
	}).First(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}
