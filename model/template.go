package model

import "github.com/factly/bindu-server/config"

// Template model
type Template struct {
	config.Base
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	URL            string `json:"url"`
	OrganisationID uint   `json:"organisation_id"`
}
