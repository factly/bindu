package model

import (
	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// Template model
type Template struct {
	config.Base
	Title      string         `json:"title"`
	Slug       string         `json:"slug"`
	Schema     postgres.Jsonb `gorm:"column:schema" json:"schema" sql:"jsonb" swaggertype:"primitive,string"`
	Properties postgres.Jsonb `gorm:"column:properties" json:"properties" sql:"jsonb" swaggertype:"primitive,string"`
	MediumID   *uint          `gorm:"column:medium_id;default:NULL" json:"medium_id"`
	Medium     *Medium        `gorm:"foreignKey:medium_id" json:"medium"`
	SpaceID    uint           `gorm:"column:space_id" json:"space_id"`
	Space      *Space         `gorm:"foreignKey:space_id" json:"space,omitempty"`
}

var templateUser config.ContextKey = "template_user"

// BeforeCreate hook
func (t *Template) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(templateUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	t.CreatedByID = uint(uID)
	t.UpdatedByID = uint(uID)
	return nil
}
