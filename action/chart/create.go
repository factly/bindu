package chart

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/factly/bindu-server/util/minio"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/slugx"
	"github.com/factly/x/validationx"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// create - Create chart
// @Summary Create chart
// @Description Create chart
// @Tags Chart
// @ID add-chart
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param Chart body chart true "Chart Object"
// @Success 201 {object} model.Chart
// @Failure 400 {array} string
// @Router /charts [post]
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

	if viper.GetBool("create_super_organisation") {
		// Fetch space permissions
		permission := model.SpacePermission{}
		err = config.DB.Model(&model.SpacePermission{}).Where(&model.SpacePermission{
			SpaceID: uint(sID),
		}).First(&permission).Error

		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot create more charts", http.StatusUnprocessableEntity)))
			return
		}

		// Fetch total number of charts in space
		var totCharts int64
		config.DB.Model(&model.Chart{}).Where(&model.Chart{
			SpaceID: uint(sID),
		}).Count(&totCharts)

		if totCharts >= permission.Charts && permission.Charts > 0 {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot create more charts", http.StatusUnprocessableEntity)))
			return
		}
	}

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
		Base: config.Base{
			CreatedByID: uint(uID),
			UpdatedByID: uint(uID),
		},
		URL:     msgJSONB,
		SpaceID: uint(sID),
	}

	tx := config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Begin()

	err = tx.Model(&model.Medium{}).Create(&medium).Error

	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	// Get table name
	stmt := &gorm.Statement{DB: config.DB}
	_ = stmt.Parse(&model.Chart{})
	tableName := stmt.Schema.Table

	var chartSlug string
	if chart.Slug != "" && slugx.Check(chart.Slug) {
		chartSlug = chart.Slug
	} else {
		chartSlug = slugx.Make(chart.Title)
	}

	themeID := &chart.ThemeID
	if chart.ThemeID == 0 {
		themeID = nil
	}

	result := &model.Chart{
		ID:               strings.ReplaceAll(uuid.New().String(), "-", ""),
		Title:            chart.Title,
		Slug:             slugx.Approve(&config.DB, chartSlug, sID, tableName),
		DataURL:          chart.DataURL,
		Config:           chart.Config,
		Description:      chart.Description,
		Status:           chart.Status,
		IsPublic:         chart.IsPublic,
		FeaturedMediumID: &medium.ID,
		ThemeID:          themeID,
		TemplateID:       chart.TemplateID,
		PublishedDate:    chart.PublishedDate,
		SpaceID:          uint(sID),
	}

	if len(chart.TagIDs) > 0 {
		tx.Model(&model.Tag{}).Where(chart.TagIDs).Find(&result.Tags)
	}
	if len(chart.CategoryIDs) > 0 {
		tx.Model(&model.Category{}).Where(chart.CategoryIDs).Find(&result.Categories)
	}

	err = tx.Model(&model.Chart{}).Create(&result).Error

	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	tx.Model(&model.Chart{}).Preload("Medium").Preload("Theme").Preload("Tags").Preload("Categories").First(&result)

	tx.Commit()

	renderx.JSON(w, http.StatusCreated, result)
}
