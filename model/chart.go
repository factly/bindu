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
	Slug             string         `json:"slug"`
	Description      postgres.Jsonb `json:"description" sql:"jsonb"`
	DataURL          string         `json:"data_url"`
	Config           postgres.Jsonb `json:"config"`
	Status           string         `json:"status"`
	FeaturedMediumID uint           `gorm:"column:featured_medium_id" json:"featured_medium_id" sql:"DEFAULT:NULL"`
	Medium           *Medium        `gorm:"foreignkey:featured_medium_id;association_foreignkey:id" json:"medium"`
	ThemeID          uint           `gorm:"column:theme_id" json:"theme_id"`
	Theme            Theme          `gorm:"foreignkey:theme_id;association_foreignkey:id" json:"theme"`
	PublishedDate    time.Time      `json:"published_date"`
	OrganisationID   uint           `json:"organisation_id"`
	Tags             []Tag          `gorm:"many2many:chart_tag;" json:"tags"`
	Categories       []Category     `gorm:"many2many:chart_category;" json:"categories"`
}
