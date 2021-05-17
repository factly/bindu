package template

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"

	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// delete - Delete template by id
// @Summary Delete a template
// @Description Delete template by ID
// @Tags Template
// @ID delete-template-by-id
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param template_id path string true "Template ID"
// @Success 200
// @Failure 400 {array} string
// @Router  /templates/{template_id} [delete]
func delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "template_id")
	if id == "" {
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	result := &model.Template{}
	result.ID = id

	// check record exists or not
	err = config.DB.Where(&model.Template{
		SpaceID: uint(sID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	if result.IsDefault {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot delete default template", http.StatusUnprocessableEntity)))
		return
	}

	config.DB.Delete(&result)

	renderx.JSON(w, http.StatusOK, nil)
}
