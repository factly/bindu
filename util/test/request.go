package test

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
)

// CleanTables - to clean tables in DB
func CleanTables() {
	config.DB.Model(&model.Chart{}).RemoveForeignKey("featured_medium_id", "bi_medium(id)")
	config.DB.Model(&model.Chart{}).RemoveForeignKey("theme_id", "bi_theme(id)")

	config.DB.DropTable("bi_chart_category")
	config.DB.DropTable("bi_chart_tag")
	config.DB.DropTable(&model.Medium{})
	config.DB.DropTable(&model.Theme{})
	config.DB.DropTable(&model.Tag{})
	config.DB.DropTable(&model.Category{})
	config.DB.DropTable(&model.Chart{})
}
