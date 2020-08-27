package model

import "github.com/factly/bindu-server/config"

// Migration does database migrations
func Migration() {
	// db migrations
	config.DB.AutoMigrate(
		&Category{},
		&Chart{},
		&Medium{},
		&Tag{},
		&Theme{},
	)

	// Adding foreignKey
	config.DB.Model(&Chart{}).AddForeignKey("featured_medium_id", "bi_medium(id)", "RESTRICT", "RESTRICT")
	config.DB.Model(&Chart{}).AddForeignKey("theme_id", "bi_theme(id)", "RESTRICT", "RESTRICT")
}
