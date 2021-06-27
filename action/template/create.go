package template

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/meilisearchx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// create - Create template
// @Summary Create template
// @Description Create template
// @Tags Template
// @ID add-template
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param Template body template true "Template Object"
// @Success 201 {object} model.Template
// @Failure 400 {array} string
// @Router /templates [post]
func create(w http.ResponseWriter, r *http.Request) {

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

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Template{})
	tableName := stmt.Schema.Table

	var templateSlug string
	if template.Slug != "" && slugx.Check(template.Slug) {
		templateSlug = template.Slug
	} else {
		templateSlug = slugx.Make(template.Title)
	}

	mediumID := &template.MediumID
	if template.MediumID == 0 {
		mediumID = nil
	}

	// Store HTML description
	var description string
	if len(template.Description.RawMessage) > 0 && !reflect.DeepEqual(template.Description, util.NilJsonb()) {
		description, err = util.HTMLDescription(template.Description)
		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot parse template description", http.StatusUnprocessableEntity)))
			return
		}
	}

	result := &model.Template{
		ID:              strings.ReplaceAll(uuid.New().String(), "-", ""),
		Title:           template.Title,
		Slug:            slugx.Approve(&config.DB, templateSlug, sID, tableName),
		Spec:            template.Spec,
		Properties:      template.Properties,
		SpaceID:         uint(sID),
		MediumID:        mediumID,
		Mode:            template.Mode,
		Description:     template.Description,
		HtmlDescription: description,
		CategoryID:      template.CategoryID,
	}

	tx := config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Begin()
	err = tx.Model(&model.Template{}).Create(&result).Error
	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	tx.Model(&result).Preload("Medium").Preload("Category").Find(&result)

	if err = AddToMeili(result); err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	tx.Commit()
	renderx.JSON(w, http.StatusCreated, result)
}

func AddToMeili(result *model.Template) error {
	// Insert into meili index
	meiliObj := map[string]interface{}{
		"id":               result.ID,
		"kind":             "template",
		"title":            result.Title,
		"slug":             result.Slug,
		"category_id":      result.CategoryID,
		"medium_id":        result.MediumID,
		"is_default":       result.IsDefault,
		"mode":             result.Mode,
		"description":      result.Description,
		"html_description": result.HtmlDescription,
		"space_id":         result.SpaceID,
	}

	err := meilisearchx.AddDocument("bindu", meiliObj)
	if err != nil {
		return err
	}
	return err
}
