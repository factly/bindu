package model

import (
	"github.com/factly/bindu-server/config"
	"gorm.io/gorm"
)

// Category model
type Category struct {
	config.Base
	Name          string   `json:"name" validate:"required"`
	Slug          string   `json:"slug" validate:"required"`
	Description   string   `json:"description"`
	IsForTemplate bool     `gorm:"column:is_for_template" json:"is_for_template"`
	SpaceID       uint     `gorm:"column:space_id" json:"space_id"`
	Space         *Space   `json:"space,omitempty"`
	Charts        []*Chart `gorm:"many2many:chart_category;" json:"charts"`
}

var categoryUser config.ContextKey = "category_user"

// BeforeCreate hook
func (category *Category) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(categoryUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	category.CreatedByID = uint(uID)
	category.UpdatedByID = uint(uID)
	return nil
}
