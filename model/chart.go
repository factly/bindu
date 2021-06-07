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
	ID               string          `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	DeletedAt        *gorm.DeletedAt `sql:"index" json:"deleted_at" swaggertype:"primitive,string"`
	CreatedByID      uint            `gorm:"column:created_by_id" json:"created_by_id"`
	UpdatedByID      uint            `gorm:"column:updated_by_id" json:"updated_by_id"`
	Title            string          `json:"title"`
	Slug             string          `json:"slug"`
	Description      postgres.Jsonb  `json:"description" sql:"jsonb" swaggertype:"primitive,string"`
	DataURL          string          `json:"data_url"`
	Config           postgres.Jsonb  `json:"config" swaggertype:"primitive,string"`
	Status           string          `json:"status"`
	IsPublic         bool            `gorm:"column:is_public" json:"is_public"`
	FeaturedMediumID *uint           `gorm:"column:featured_medium_id;default:NULL" json:"featured_medium_id"`
	Medium           *Medium         `gorm:"foreignKey:featured_medium_id" json:"medium"`
	TemplateID       string          `gorm:"column:template_id" json:"template_id"`
	Template         *Template       `gorm:"foreignKey:template_id;default:NULL" json:"template"`
	ThemeID          *uint           `gorm:"column:theme_id" json:"theme_id"`
	Theme            *Theme          `gorm:"foreignKey:theme_id;default:NULL" json:"theme"`
	PublishedDate    *time.Time      `gorm:"column:published_date" sql:"DEFAULT:NULL" json:"published_date"`
	Mode             string          `gorm:"column:mode" json:"mode"`
	SpaceID          uint            `gorm:"column:space_id" json:"space_id"`
	Space            *Space          `json:"space,omitempty"`
	Tags             []Tag           `gorm:"many2many:chart_tag;" json:"tags"`
	Categories       []Category      `gorm:"many2many:chart_category;" json:"categories"`
}

// BeforeSave - to check space for medium & theme
func (c *Chart) BeforeSave(tx *gorm.DB) (e error) {

	if c.FeaturedMediumID != nil && *c.FeaturedMediumID > 0 {
		medium := Medium{}
		medium.ID = *c.FeaturedMediumID

		err := tx.Model(&Medium{}).Where(Medium{
			SpaceID: c.SpaceID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("medium do not belong to same space")
		}
	}

	if c.ThemeID != nil && *c.ThemeID > 0 {
		theme := Theme{}
		theme.ID = *c.ThemeID

		err := tx.Model(&Theme{}).Where(Theme{
			SpaceID: c.SpaceID,
		}).First(&theme).Error

		if err != nil {
			return errors.New("theme do not belong to same space")
		}
	}

	for _, tag := range c.Tags {
		if tag.SpaceID != c.SpaceID {
			return errors.New("some tags do not belong to same space")
		}
	}

	for _, category := range c.Categories {
		if category.SpaceID != c.SpaceID {
			return errors.New("some categories do not belong to same space")
		}
	}

	return nil
}

var chartUser config.ContextKey = "chart_user"

// BeforeCreate hook
func (c *Chart) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(chartUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	c.CreatedByID = uint(uID)
	c.UpdatedByID = uint(uID)
	return nil
}
