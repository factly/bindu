package template

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
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

	result := &model.Template{
		Title:      template.Title,
		Slug:       slugx.Approve(&config.DB, templateSlug, sID, tableName),
		Schema:     template.Schema,
		Properties: template.Properties,
		SpaceID:    uint(sID),
		MediumID:   mediumID,
	}

	err = config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Model(&model.Template{}).Preload("Medium").Create(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
