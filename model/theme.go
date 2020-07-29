package model

import "github.com/factly/bindu-server/config"

// Theme model
type Theme struct {
	config.Base
	Name           string `json:"name" validate:"required"`
	Slug           string `json:"slug" validate:"required"`
	Description    string `json:"description"`
	URL            string `json:"url" validate:"required"`
	OrganisationID uint   `json:"organisation_id"`
}
