package model

import (
	"errors"
	"time"

	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
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
	FeaturedMediumID uint           `gorm:"column:featured_medium_id;default:NULL" json:"featured_medium_id"`
	Medium           *Medium        `gorm:"foreignKey:featured_medium_id" json:"medium"`
	ThemeID          uint           `gorm:"column:theme_id" json:"theme_id"`
	Theme            *Theme         `gorm:"foreignKey:theme_id;default:NULL" json:"theme"`
	PublishedDate    time.Time      `json:"published_date"`
	OrganisationID   uint           `json:"organisation_id"`
	Tags             []Tag          `gorm:"many2many:chart_tag;" json:"tags"`
	Categories       []Category     `gorm:"many2many:chart_category;" json:"categories"`
}

// BeforeSave - to check organisation for medium & theme
func (c *Chart) BeforeSave(tx *gorm.DB) (e error) {

	if c.FeaturedMediumID > 0 {
		medium := Medium{}
		medium.ID = c.FeaturedMediumID

		err := tx.Model(&Medium{}).Where(Medium{
			OrganisationID: c.OrganisationID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("medium do not belong to same organisation")
		}
	}

	if c.ThemeID > 0 {
		theme := Theme{}
		theme.ID = c.ThemeID

		err := tx.Model(&Theme{}).Where(Theme{
			OrganisationID: c.OrganisationID,
		}).First(&theme).Error

		if err != nil {
			return errors.New("theme do not belong to same organisation")
		}
	}

	for _, tag := range c.Tags {
		if tag.OrganisationID != c.OrganisationID {
			return errors.New("some tags do not belong to same organisation")
		}
	}

	for _, category := range c.Categories {
		if category.OrganisationID != c.OrganisationID {
			return errors.New("some categories do not belong to same organisation")
		}
	}

	return nil
}
