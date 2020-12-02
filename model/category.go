package model

import (
	"github.com/factly/bindu-server/config"
	"gorm.io/gorm"
)

// Category model
type Category struct {
	config.Base
	Name           string   `json:"name" validate:"required"`
	Slug           string   `json:"slug" validate:"required"`
	Description    string   `json:"description"`
	OrganisationID uint     `json:"organisation_id"`
	Charts         []*Chart `gorm:"many2many:chart_category;" json:"charts"`
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
