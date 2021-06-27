package model

import (
	"github.com/factly/bindu-server/config"
	"gorm.io/gorm"
)

type organisationUser struct {
	config.Base
	Role string `json:"role"`
}

// Organisation model
type Organisation struct {
	config.Base
	Title       string           `json:"title"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Permission  organisationUser `json:"permission"`
}

// OrganisationPermission model
type OrganisationPermission struct {
	config.Base
	OrganisationID uint  `gorm:"column:organisation_id" json:"organisation_id"`
	Spaces         int64 `gorm:"column:spaces" json:"spaces"`
}

var organisationPermissionUser config.ContextKey = "org_perm_user"

// BeforeCreate hook
func (op *OrganisationPermission) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(organisationPermissionUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	op.CreatedByID = uint(uID)
	op.UpdatedByID = uint(uID)
	return nil
}
