package template

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/renderx"
)

type categoryTemplate struct {
	model.Category
	Templates []model.Template `json:"templates"`
}

// list - Get all templates
// @Summary Show all templates
// @Description Get all templates
// @Tags Template
// @ID get-all-templates
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {array} []categoryTemplate
// @Router /templates [get]
func list(w http.ResponseWriter, r *http.Request) {

	result := make([]categoryTemplate, 0)
	templates := make([]model.Template, 0)

	config.DB.Preload("Medium").Preload("Category").Model(&model.Template{}).Order("id desc").Find(&templates)

	for _, template := range templates {
		exists := false
		for index, category := range result {
			if category.ID == template.CategoryID {
				exists = true
				category.Templates = append(category.Templates, template)
				result[index] = category
				break
			}
		}
		if exists == false {
			category := categoryTemplate{
				Category: template.Category,
			}
			category.Templates = append(category.Templates, template)
			result = append(result, category)
		}
	}

	renderx.JSON(w, http.StatusOK, result)
}
