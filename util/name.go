package util

import (
	"strings"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
)

// CheckName checks if the table contains any entry with same name
func CheckName(space uint, name, table string) bool {
	var count int64
	newName := strings.ToLower(strings.TrimSpace(name))
	config.DB.Table(table).Where("deleted_at IS NULL AND (space_id = ? AND name ILIKE ?)", space, newName).Count(&count)
	return count > 0
}

// CheckCategoryName checks if the category table contains any entry with same name
func CheckCategoryName(space uint, name string, IsForTemplate bool) bool {
	var count int64
	newName := strings.ToLower(strings.TrimSpace(name))
	config.DB.Model(&model.Category{}).Where("deleted_at IS NULL AND (space_id = ? AND name ILIKE ? AND is_for_template = ?)", space, newName, IsForTemplate).Count(&count)
	return count > 0
}
