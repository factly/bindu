package test

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
)

// Init - Initialize test db
func Init() {
	config.SetupDB("postgres://postgres:postgres@localhost:5432/bindu-test?sslmode=disable")

	// db migrations
	config.DB.AutoMigrate(
		&model.Category{},
		&model.Chart{},
		&model.Medium{},
		&model.Tag{},
		&model.Theme{},
	)

	// Adding foreignKey
	config.DB.Model(&model.Chart{}).AddForeignKey("featured_medium_id", "bi_medium(id)", "RESTRICT", "RESTRICT")
	config.DB.Model(&model.Chart{}).AddForeignKey("theme_id", "bi_theme(id)", "RESTRICT", "RESTRICT")

}
