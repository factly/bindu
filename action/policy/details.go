package policy

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/factly/bindu-server/action/user"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// details - Get policy by ID
// @Summary Get policy by ID
// @Description Get policy by ID
// @Tags Policy
// @ID get-policy-by-id
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param policy_id path string true "Policy ID"
// @Success 200 {object} model.Policy
// @Router /policies/{policy_id} [get]
func details(w http.ResponseWriter, r *http.Request) {
	spaceID, err := util.GetSpace(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	userID, err := util.GetUser(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	organisationID, err := util.GetOrganisation(r.Context())

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	policyID := chi.URLParam(r, "policy_id")

	ketoPolicyID := fmt.Sprint("id:org:", organisationID, ":app:bindu:space:", spaceID, ":", policyID)

	req, err := http.NewRequest("GET", viper.GetString("keto_url")+"/engines/acp/ory/regex/policies/"+ketoPolicyID, nil)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.NetworkError()))
		return
	}

	defer resp.Body.Close()

	ketoPolicy := model.KetoPolicy{}
	err = json.NewDecoder(resp.Body).Decode(&ketoPolicy)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	/* User req */
	userMap := user.Mapper(organisationID, userID)

	result := Mapper(ketoPolicy, userMap)

	renderx.JSON(w, http.StatusOK, result)
}
