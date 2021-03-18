package model

import (
	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// Medium model
type Medium struct {
	config.Base
	Name        string         `gorm:"column:name" json:"name"`
	Slug        string         `gorm:"column:slug" json:"slug"`
	Type        string         `gorm:"column:type" json:"type"`
	Title       string         `gorm:"column:title" json:"title"`
	Description string         `gorm:"column:description" json:"description"`
	Caption     string         `gorm:"column:caption" json:"caption"`
	AltText     string         `gorm:"column:alt_text" json:"alt_text"`
	FileSize    int64          `gorm:"column:file_size" json:"file_size"`
	URL         postgres.Jsonb `gorm:"column:url" json:"url" swaggertype:"primitive,string"`
	Dimensions  string         `gorm:"column:dimensions" json:"dimensions"`
	SpaceID     uint           `gorm:"column:space_id" json:"space_id"`
}

var mediumUser config.ContextKey = "medium_user"

// BeforeCreate hook
func (media *Medium) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(mediumUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	media.CreatedByID = uint(uID)
	media.UpdatedByID = uint(uID)
	return nil
}
