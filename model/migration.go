package model

import "github.com/factly/bindu-server/config"

// Migration does database migrations
func Migration() {
	// db migrations
	_ = config.DB.AutoMigrate(
		&Medium{},
		&Chart{},
		&Category{},
		&Tag{},
		&Theme{},
		&Template{},
		&OrganisationPermission{},
		&SpacePermission{},
		&OrganisationPermissionRequest{},
		&SpacePermissionRequest{},
	)
}
