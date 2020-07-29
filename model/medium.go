package model

import (
	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Medium model
type Medium struct {
	config.Base
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	Type           string         `json:"type"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Caption        string         `json:"caption"`
	AltText        string         `json:"alt_text"`
	FileSize       int64          `json:"file_size"`
	URL            postgres.Jsonb `json:"url"`
	Dimensions     string         `json:"dimensions"`
	OrganisationID uint           `json:"organisation_id"`
}
