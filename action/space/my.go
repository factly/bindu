package space

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"

	"github.com/factly/bindu-server/model"
)

type orgWithSpace struct {
	model.Organisation
	Spaces []model.Space `json:"spaces"`
}

// list - Get all spaces for a user
// @Summary Show all spaces
// @Description Get all spaces
// @Tags Space
// @ID get-all-spaces
// @Produce  json
// @Param X-User header string true "User ID"
// @Success 200 {array} []orgWithSpace
// @Router /spaces [get]
func my(w http.ResponseWriter, r *http.Request) {
	orgList, err := util.RequestOrganisation(r)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	var allOrgIDs []int

	for _, each := range orgList {
		allOrgIDs = append(allOrgIDs, int(each.ID))
	}

	// Fetched all the spaces related to all the organisations
	var allSpaces = make([]model.Space, 0)

	config.DB.Model(model.Space{}).Where("organisation_id IN (?)", allOrgIDs).Preload("Logo").Preload("LogoMobile").Preload("FavIcon").Preload("MobileIcon").Find(&allSpaces)

	result := make([]orgWithSpace, 0)

	for _, organisation := range orgList {
		spaces := make([]model.Space, 0)
		for _, space := range allSpaces {
			if space.OrganisationID == int(space.OrganisationID) {
				spaces = append(spaces, space)
			}
		}

		os := orgWithSpace{
			Organisation: organisation,
			Spaces:       spaces,
		}
		result = append(result, os)
	}

	renderx.JSON(w, http.StatusOK, result)
}
