package policy

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/factly/bindu-server/util"

	"github.com/spf13/viper"

	"github.com/factly/bindu-server/model"
	"github.com/factly/x/loggerx"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Composer create keto policy
func Composer(oID int, sID int, inputPolicy policyReq) model.KetoPolicy {
	allowedResources := []string{"categories", "charts", "media", "policies", "roles", "tags", "themes"}
	allowedActions := []string{"get", "create", "update", "delete", "publish"}
	result := model.KetoPolicy{}

	commanPolicyString := fmt.Sprint(":org:", oID, ":app:bindu:space:", sID, ":")
	result.ID = "id" + commanPolicyString + inputPolicy.Name
	result.Description = inputPolicy.Description
	result.Effect = "allow"
	result.Resources = make([]string, 0)
	result.Actions = make([]string, 0)

	for _, each := range inputPolicy.Permissions {
		if contains(allowedResources, each.Resource) {
			result.Resources = append(result.Resources, "resources"+commanPolicyString+each.Resource)
			var eachActions []string
			for _, action := range each.Actions {
				if contains(allowedActions, action) {
					eachActions = append(eachActions, "actions"+commanPolicyString+each.Resource+":"+action)
				}
			}
			result.Actions = append(result.Actions, eachActions...)
		}
	}

	inpSubjects := make([]string, 0)

	for _, subject := range inputPolicy.Subjects {
		_, err := strconv.Atoi(subject)
		if err != nil {
			inpSubjects = append(inpSubjects, fmt.Sprint("roles:org:", oID, ":app:bindu:space:", sID, ":", subject))
		} else {
			inpSubjects = append(inpSubjects, subject)
		}
	}

	result.Subjects = inpSubjects

	resp, err := util.Request("PUT", viper.GetString("keto_url")+"/engines/acp/ory/regex/policies", result)
	if err != nil {
		return model.KetoPolicy{}
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		loggerx.Error(err)
	}

	return result
}
