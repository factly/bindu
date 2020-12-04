package model

import (
	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// Theme model
type Theme struct {
	config.Base
	Name    string         `json:"name" validate:"required"`
	Config  postgres.Jsonb `json:"config" swaggertype:"primitive,string"`
	SpaceID uint           `gorm:"column:space_id" json:"space_id"`
	Space   *Space         `json:"space,omitempty"`
}

var themeUser config.ContextKey = "theme_user"

// BeforeCreate hook
func (theme *Theme) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(themeUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	theme.CreatedByID = uint(uID)
	theme.UpdatedByID = uint(uID)
	return nil
}
