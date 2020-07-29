package model

import "github.com/factly/bindu-server/config"

// Tag model
type Tag struct {
	config.Base
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	OrganisationID uint   `json:"organisation_id"`
}
