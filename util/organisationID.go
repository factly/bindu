package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/factly/x/errorx"
)

type ctxKeyOrganisationID int

// OrganisationIDKey is the key that holds the unique organisation ID in a request context.
const OrganisationIDKey ctxKeyOrganisationID = 0

// GenerateOrganisation check X-Organisation in header
func GenerateOrganisation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Orgs")
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

		orgs := RequestOrganisation(w, r)
		foundOrg := false

		for _, each := range orgs {
			eachOrg := (each).(map[string]interface{})
			if int((eachOrg["id"]).(float64)) == oID {
				foundOrg = true
				break
			}
		}

		if foundOrg == false {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, OrganisationIDKey, oID)
		h.ServeHTTP(w, r.WithContext(ctx))
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
func RequestOrganisation(w http.ResponseWriter, r *http.Request) []interface{} {

	uID, err := GetUser(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return nil
	}

	req, err := http.NewRequest("GET", os.Getenv("KAVACH_URL")+"/organisations/my", nil)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User", fmt.Sprint(uID))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.NetworkError()))
		return nil
	}

	defer resp.Body.Close()

	var orgs []interface{}
	json.NewDecoder(resp.Body).Decode(&orgs)

	return orgs
}
