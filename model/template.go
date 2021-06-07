package model

import (
	"time"

	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// Template model
type Template struct {
	ID          string          `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `sql:"index" json:"deleted_at" swaggertype:"primitive,string"`
	CreatedByID uint            `gorm:"column:created_by_id" json:"created_by_id"`
	UpdatedByID uint            `gorm:"column:updated_by_id" json:"updated_by_id"`
	Title       string          `json:"title"`
	Slug        string          `json:"slug"`
	Spec        postgres.Jsonb  `gorm:"column:spec" json:"spec" sql:"jsonb" swaggertype:"primitive,string"`
	Properties  postgres.Jsonb  `gorm:"column:properties" json:"properties" sql:"jsonb" swaggertype:"primitive,string"`
	CategoryID  uint            `gorm:"column:category_id" json:"category_id"`
	Category    Category        `gorm:"foreignKey:category_id" json:"category"`
	MediumID    *uint           `gorm:"column:medium_id;default:NULL" json:"medium_id"`
	Medium      *Medium         `gorm:"foreignKey:medium_id" json:"medium"`
	IsDefault   bool            `gorm:"column:is_default" json:"is_default"`
	Mode        string          `gorm:"column:mode" json:"mode"`
	SpaceID     uint            `gorm:"column:space_id" json:"space_id"`
	Space       *Space          `gorm:"foreignKey:space_id" json:"space,omitempty"`
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
