package role

import (
	"strings"

	"github.com/factly/bindu-server/model"
)

// Mapper created model.Role object from keto role
func Mapper(ketoRole model.KetoRole, userMap map[string]model.User) model.Role {
	result := model.Role{}
	result.Description = ketoRole.Description

	roleIDArr := strings.Split(ketoRole.ID, ":")
	result.Name = roleIDArr[len(roleIDArr)-1]

	users := make([]model.User, 0)
	for _, user := range ketoRole.Members {
		val, exists := userMap[user]
		if exists {
			users = append(users, val)
		}
	}

	result.Users = users
	return result
}
