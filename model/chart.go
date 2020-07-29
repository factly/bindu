package model

import (
	"time"

	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Chart model
type Chart struct {
	config.Base
	Title            string         `json:"title"`
	Subtitle         string         `json:"subtitle"`
	Slug             string         `json:"slug"`
	URL              string         `json:"url"`
	Description      postgres.Jsonb `json:"description" sql:"jsonb"`
	Status           string         `json:"status"`
	FeaturedMediumID uint           `gorm:"column:featured_medium_id" json:"featured_medium_id" sql:"DEFAULT:NULL"`
	Medium           *Medium        `gorm:"foreignkey:featured_medium_id;association_foreignkey:id" json:"medium"`
	TemplateID       uint           `gorm:"column:template_id" json:"template_id"`
	Template         Template       `gorm:"foreignkey:template_id;association_foreignkey:id" json:"template"`
	ThemeID          uint           `gorm:"column:theme_id" json:"theme_id"`
	Theme            Theme          `gorm:"foreignkey:theme_id;association_foreignkey:id" json:"theme"`
	PublishedDate    time.Time      `json:"published_date"`
	OrganisationID   uint           `json:"organisation_id"`
}
