package user

import (
	"context"

	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/middlewarex"
)

// All - to return all authors
func All(ctx context.Context) (map[string]model.User, error) {
	authors := make(map[string]model.User)

	organisationID, err := util.GetOrganisation(ctx)

	if err != nil {
		return authors, err
	}

	userID, err := middlewarex.GetUser(ctx)

	if err != nil {
		return authors, err
	}

	authors = Mapper(organisationID, userID)

	return authors, nil

}
