package role

import (
	"encoding/json"
	"fmt"

	"github.com/factly/bindu-server/util"
	"github.com/spf13/viper"

	"github.com/factly/bindu-server/model"
)

// Composer created role in keto
func Composer(oID, sID int, inpRole roleReq) (model.KetoRole, error) {
	role := model.KetoRole{}
	role.ID = fmt.Sprint("roles:org:", oID, ":app:bindu:space:", sID, ":", inpRole.Name)
	role.Description = inpRole.Description
	role.Members = inpRole.Users

	res, err := util.Request("PUT", viper.GetString("keto_url")+"/engines/acp/ory/regex/roles", role)
	if err != nil {
		return model.KetoRole{}, err
	}

	var result model.KetoRole
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return model.KetoRole{}, err
	}

	return result, nil
}
