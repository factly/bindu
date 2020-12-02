package model

import (
	"github.com/factly/bindu-server/config"
	"gorm.io/gorm"
)

// Tag model
type Tag struct {
	config.Base
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Description    string   `json:"description"`
	OrganisationID uint     `json:"organisation_id"`
	Charts         []*Chart `gorm:"many2many:chart_tag;" json:"charts"`
}

var tagUser config.ContextKey = "tag_user"

// BeforeCreate hook
func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(tagUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	tag.CreatedByID = uint(uID)
	tag.UpdatedByID = uint(uID)
	return nil
}
