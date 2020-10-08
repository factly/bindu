package util

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/factly/bindu-server/model"
	"github.com/spf13/viper"
)

type ctxKeyOrganisationID int

// OrganisationIDKey is the key that holds the unique organisation ID in a request context.
const OrganisationIDKey ctxKeyOrganisationID = 0

// CheckOrganisation check X-Organisation in header
func CheckOrganisation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Trim(r.URL.Path, "/") != "organisations" {
			org := r.Header.Get("X-Organisation")
			if org == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			oID, err := strconv.Atoi(org)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			orgs, err := RequestOrganisation(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			foundOrg := false
			for _, each := range orgs {
				if each.Base.ID == uint(oID) {
					foundOrg = true
					break
				}
			}

			if !foundOrg {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, OrganisationIDKey, oID)
			h.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		h.ServeHTTP(w, r)
	})
}

// GetOrganisation return organisation ID
func GetOrganisation(ctx context.Context) (int, error) {
	if ctx == nil {
		return 0, errors.New("context not found")
	}
	organisationID := ctx.Value(OrganisationIDKey)
	if organisationID != nil {
		return organisationID.(int), nil
	}
	return 0, errors.New("something went wrong")
}

// RequestOrganisation - request kavach to get all organisations of user
func RequestOrganisation(r *http.Request) ([]model.Organisation, error) {

	uID, err := GetUser(r.Context())

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", viper.GetString("kavach.url")+"/organisations/my", nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("X-User", strconv.Itoa(uID))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	orgs := []model.Organisation{}
	err = json.Unmarshal(body, &orgs)

	return orgs, err
}
