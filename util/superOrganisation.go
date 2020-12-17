package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/factly/bindu-server/model"
	"github.com/spf13/viper"
)

// CheckSuperOrganisation checks weather organisation of user is super org or not
func CheckSuperOrganisation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !viper.GetBool("create_super_organisation") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oID, err := GetOrganisation(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		superOrgID, err := GetSuperOrganisationID()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if oID != superOrgID {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// GetSuperOrganisationID get superorganisation id from keto policy
func GetSuperOrganisationID() (int, error) {
	resp, err := Request("GET", viper.GetString("keto_url")+"/engines/acp/ory/regex/policies/app:bindu:superorg", nil)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode == http.StatusOK {
		var policy model.KetoPolicy
		err = json.NewDecoder(resp.Body).Decode(&policy)
		if err != nil {
			return 0, err
		}

		if len(policy.Subjects) != 0 {
			orgID, _ := strconv.Atoi(policy.Subjects[0])
			return orgID, nil
		}
	}
	return 0, errors.New("cannot get super organisation id")
}
