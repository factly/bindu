package policy

import (
	"encoding/json"
	"net/http"

	"github.com/factly/bindu-server/action/user"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/go-chi/chi"
)

// update - Update policy
// @Summary Update policy
// @Description Update policy
// @Tags Policy
// @ID update-policy
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param policy_id path string true "Policy ID"
// @Param Policy body policyReq true "Policy Object"
// @Success 200 {object} model.Policy
// @Router /policies/{policy_id} [put]
func update(w http.ResponseWriter, r *http.Request) {
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

	/* create new policy */
	policyReq := policyReq{}

	err = json.NewDecoder(r.Body).Decode(&policyReq)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	policyID := chi.URLParam(r, "policy_id")
	if policyID == "" {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InvalidID()))
		return
	}

	/* User req */
	policyReq.Name = policyID
	result := Mapper(Composer(organisationID, spaceID, policyReq), user.Mapper(organisationID, userID))

	renderx.JSON(w, http.StatusOK, result)
}
