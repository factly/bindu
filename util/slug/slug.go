package slug

import (
	"strconv"

	"github.com/factly/bindu-server/config"
)

type comman struct {
	Slug string `gorm:"column:slug" json:"slug"`
}

// Approve return slug
func Approve(slug string, organisation int, table string) string {
	var result []comman
	config.DB.Table(table).Select("slug, organisation_id").Where("slug LIKE ? AND organisation_id = ?", slug+"%", organisation).Scan(&result)
	count := 0
	for {
		flag := true
		for _, each := range result {
			temp := slug
			if count != 0 {
				temp = temp + "-" + strconv.Itoa(count)
			}
			if each.Slug == temp {
				flag = false
				break
			}
		}
		if flag {
			break
		}
		count++
	}
	temp := slug
	if count != 0 {
		temp = temp + "-" + strconv.Itoa(count)
	}
	return temp
}
