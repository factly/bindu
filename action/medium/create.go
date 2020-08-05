package medium

import (
	"encoding/json"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
)

// create - Create medium
// @Summary Create medium
// @Description Create medium
// @Tags Medium
// @ID add-medium
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param Medium body medium true "Medium Object"
// @Success 201 {object} model.Medium
// @Failure 400 {array} string
// @Router /media [post]
func create(w http.ResponseWriter, r *http.Request) {

	oID, err := util.GetOrganisation(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	medium := &medium{}

	json.NewDecoder(r.Body).Decode(&medium)

	validationError := validationx.Check(medium)

	if validationError != nil {
		errorx.Render(w, validationError)
		return
	}

	var mediumSlug string
	if medium.Slug != "" && slug.Check(medium.Slug) {
		mediumSlug = medium.Slug
	} else {
		mediumSlug = slug.Make(medium.Name)
	}

	result := &model.Medium{
		Name:           medium.Name,
		Slug:           slug.Approve(mediumSlug, oID, config.DB.NewScope(&model.Medium{}).TableName()),
		Type:           medium.Type,
		URL:            medium.URL,
		OrganisationID: uint(oID),
	}

	err = config.DB.Model(&model.Medium{}).Create(&result).Error

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
