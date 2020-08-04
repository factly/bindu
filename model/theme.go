package model

import (
	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Theme model
type Theme struct {
	config.Base
	Name           string         `json:"name" validate:"required"`
	Config         postgres.Jsonb `json:"config"`
	OrganisationID uint           `json:"organisation_id"`
}
