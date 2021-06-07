package chart

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util/minio"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/meilisearchx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// update - Update chart by id
// @Summary Update a chart by id
// @Description Update chart by ID
// @Tags Chart
// @ID update-chart-by-id
// @Produce json
// @Consume json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param chart_id path string true "Chart ID"
// @Param Chart body chart false "Chart"
// @Success 200 {object} model.Chart
// @Router /charts/{chart_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "chart_id")
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

	chart := &chart{}
	err = json.NewDecoder(r.Body).Decode(&chart)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(chart)
	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	result := &model.Chart{}
	result.ID = id

	// check record exists or not
	err = config.DB.Where(&model.Chart{
		SpaceID: uint(sID),
	}).First(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.RecordNotFound()))
		return
	}

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Chart{})
	tableName := stmt.Schema.Table

	var chartSlug string

	if result.Slug == chart.Slug {
		chartSlug = result.Slug
	} else if chart.Slug != "" && slugx.Check(chart.Slug) {
		chartSlug = slugx.Approve(&config.DB, chart.Slug, sID, tableName)
	} else {
		chartSlug = slugx.Approve(&config.DB, slugx.Make(chart.Title), sID, tableName)
	}

	tx := config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Begin()
	if chart.FeaturedMediumID == 0 {
		mediaURL, err := minio.Upload(r, chart.FeaturedMedium)

		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
			return
		}

		mediumJSON := map[string]interface{}{
			"url": mediaURL,
		}

		mediumByte, err := json.Marshal(mediumJSON)
		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
			return
		}

		msg := json.RawMessage(mediumByte)

		var msgJSONB postgres.Jsonb

		msgJSONB.RawMessage = msg

		medium := model.Medium{
			URL:     msgJSONB,
			SpaceID: uint(sID),
		}
		err = tx.Model(&model.Medium{}).Create(&medium).Error
		if err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
		chart.FeaturedMediumID = medium.ID
	}

	newTags := make([]model.Tag, 0)
	if len(chart.TagIDs) > 0 {
		tx.Model(&model.Tag{}).Where(chart.TagIDs).Find(&newTags)
		if err = tx.Model(&result).Association("Tags").Replace(&newTags); err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
	} else {
		_ = tx.Model(&result).Association("Tags").Clear()
	}

	newCategories := make([]model.Category, 0)
	if len(chart.CategoryIDs) > 0 {
		tx.Model(&model.Category{}).Where(chart.CategoryIDs).Find(&newCategories)
		if err = tx.Model(&result).Association("Categories").Replace(&newCategories); err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
	} else {
		_ = tx.Model(&result).Association("Categories").Clear()
	}

	themeID := &chart.ThemeID
	result.ThemeID = &chart.ThemeID
	if chart.ThemeID == 0 {
		err = tx.Model(&result).Omit("Tags", "Categories").Updates(map[string]interface{}{"theme_id": nil}).Error
		themeID = nil
		if err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}
	}

	tx.Model(&result).Select("IsPublic").Updates(model.Chart{IsPublic: chart.IsPublic})
	err = tx.Model(&result).Omit("Tags", "Categories").Updates(model.Chart{
		UpdatedByID:      uint(uID),
		Title:            chart.Title,
		Slug:             chartSlug,
		DataURL:          chart.DataURL,
		Description:      chart.Description,
		Status:           chart.Status,
		FeaturedMediumID: &chart.FeaturedMediumID,
		Config:           chart.Config,
		ThemeID:          themeID,
		TemplateID:       chart.TemplateID,
		PublishedDate:    chart.PublishedDate,
		Mode:             chart.Mode,
	}).Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").Preload("Template").First(&result).Error

	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	// Update into meili index
	var meiliPublishDate int64
	if result.PublishedDate != nil {
		meiliPublishDate = result.PublishedDate.Unix()
	}
	meiliObj := map[string]interface{}{
		"id":                 result.ID,
		"kind":               "chart",
		"title":              result.Title,
		"slug":               result.Slug,
		"status":             result.Status,
		"description":        result.Description,
		"data_url":           result.DataURL,
		"is_public":          result.IsPublic,
		"featured_medium_id": result.FeaturedMediumID,
		"published_date":     meiliPublishDate,
		"template_id":        result.TemplateID,
		"theme_id":           result.ThemeID,
		"space_id":           result.SpaceID,
		"tag_ids":            chart.TagIDs,
		"category_ids":       chart.CategoryIDs,
		"mode":               chart.Mode,
	}

	err = meilisearchx.UpdateDocument("bindu", meiliObj)
	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	tx.Commit()
	renderx.JSON(w, http.StatusOK, result)
}
