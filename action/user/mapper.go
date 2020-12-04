package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factly/bindu-server/model"
	"github.com/factly/x/loggerx"
	"github.com/spf13/viper"
)

// Mapper map user with id
// if any error occurs then Mapper just returns empty list
func Mapper(oID int, uID int) map[string]model.User {
	userMap := make(map[string]model.User)
	url := fmt.Sprint(viper.GetString("kavach_url"), "/organisations/", oID, "/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		loggerx.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User", fmt.Sprint(uID))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return userMap
	}

	defer resp.Body.Close()

	users := []model.User{}
	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		loggerx.Error(err)
	}

	for _, u := range users {
		userMap[fmt.Sprint(u.ID)] = u
	}

	return userMap
}
