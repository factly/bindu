package model

import "github.com/factly/bindu-server/config"

// Category model
type Category struct {
	config.Base
	Name           string   `json:"name" validate:"required"`
	Slug           string   `json:"slug" validate:"required"`
	Description    string   `json:"description"`
	OrganisationID uint     `json:"organisation_id"`
	Charts         []*Chart `gorm:"many2many:chart_category;" json:"charts"`
}
