package template

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/meilisearchx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// update - Update template by id
// @Summary Update a template by id
// @Description Update template by ID
// @Tags Template
// @ID update-template-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param template_id path string true "Template ID"
// @Param Template body template false "Template"
// @Success 200 {object} model.Template
// @Router /templates/{template_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
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

	uID, err := middlewarex.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	template := &template{}
	err = json.NewDecoder(r.Body).Decode(&template)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}
	validationError := validationx.Check(template)

	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
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
		errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot edit default template", http.StatusUnprocessableEntity)))
		return
	}

	tx := config.DB.Begin()

	mediumID := &template.MediumID
	result.MediumID = &template.MediumID
	if template.MediumID == 0 {
		err = tx.Model(&result).Updates(map[string]interface{}{"medium_id": nil}).Error
		mediumID = nil
		if err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
	}

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Template{})
	tableName := stmt.Schema.Table

	var templateSlug string

	if result.Slug == template.Slug {
		templateSlug = result.Slug
	} else if template.Slug != "" && slugx.Check(template.Slug) {
		templateSlug = slugx.Approve(&tx, template.Slug, sID, tableName)
	} else {
		templateSlug = slugx.Approve(&tx, slugx.Make(template.Title), sID, tableName)
	}

	tx.Model(&result).Updates(model.Template{
		UpdatedByID: uint(uID),
		Title:       template.Title,
		Slug:        templateSlug,
		Spec:        template.Spec,
		Properties:  template.Properties,
		MediumID:    mediumID,
		CategoryID:  template.CategoryID,
	}).Preload("Medium").Preload("Category").First(&result)

	if err = UpdateInMeili(result); err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	tx.Commit()

	renderx.JSON(w, http.StatusOK, result)
}

func UpdateInMeili(result *model.Template) error {
	// Insert into meili index
	meiliObj := map[string]interface{}{
		"id":          result.ID,
		"kind":        "template",
		"title":       result.Title,
		"slug":        result.Slug,
		"category_id": result.CategoryID,
		"medium_id":   result.MediumID,
		"is_default":  result.IsDefault,
		"space_id":    result.SpaceID,
	}

	err := meilisearchx.UpdateDocument("bindu", meiliObj)
	if err != nil {
		return err
	}
	return err
}
