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
	URL            postgres.Jsonb `json:"url" swaggertype:"primitive,string"`
	OrganisationID uint           `json:"organisation_id"`
}
