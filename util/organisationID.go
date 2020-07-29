package util

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type ctxKeyOrganisationID int

// OrganisationIDKey is the key that holds the unique organisation ID in a request context.
const OrganisationIDKey ctxKeyOrganisationID = 0

// GenerateOrganisation check X-Organisation in header
func GenerateOrganisation(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		org := r.Header.Get("X-Organisation")
		if org == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		uid, err := strconv.Atoi(org)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, OrganisationIDKey, uid)
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
