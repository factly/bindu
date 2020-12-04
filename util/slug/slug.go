package slug

import (
	"strconv"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
)

type comman struct {
	Slug string `gorm:"column:slug" json:"slug"`
}

// Approve return slug
func Approve(slug string, space int, table string) string {
	var result []comman
	config.DB.Table(table).Select("slug, space_id").Where("slug LIKE ? AND space_id = ? AND deleted_at IS NULL", slug+"%", space).First(&result)
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

// ApproveSpaceSlug return slug for space
func ApproveSpaceSlug(slug string) string {
	spaceList := make([]model.Space, 0)
	config.DB.Model(&model.Space{}).Where("slug LIKE ? AND deleted_at IS NULL", slug+"%").Find(&spaceList)

	count := 0
	for {
		flag := true
		for _, each := range spaceList {
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
