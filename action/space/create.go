package space

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/slug"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
	"github.com/spf13/viper"
)

// create - Create space
// @Summary Create space
// @Description Create space
// @Tags Space
// @ID add-space
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param Space body space true "Space Object"
// @Success 201 {object} model.Space
// @Router /spaces [post]
func create(w http.ResponseWriter, r *http.Request) {
	uID, err := util.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	space := &space{}

	err = json.NewDecoder(r.Body).Decode(&space)

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(space)

	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	if space.OrganisationID == 0 {
		return
	}

	// Check if the logger in user is admin of organisation
	err = util.CheckSpaceKetoPermission("create", uint(space.OrganisationID), uint(uID))
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.GetMessage(err.Error(), http.StatusUnauthorized)))
		return
	}

	var superOrgID int
	if viper.GetBool("create_super_organisation") {
		superOrgID, err = util.GetSuperOrganisationID()
		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
			return
		}

		// Fetch organisation permissions
		permission := model.OrganisationPermission{}
		err = config.DB.Model(&model.OrganisationPermission{}).Where(&model.OrganisationPermission{
			OrganisationID: uint(space.OrganisationID),
		}).First(&permission).Error

		if err != nil && space.OrganisationID != superOrgID {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot create more spaces", http.StatusUnprocessableEntity)))
			return
		}

		if err == nil {
			// Fetch total number of spaces in organisation
			var totSpaces int64
			config.DB.Model(&model.Space{}).Where(&model.Space{
				OrganisationID: space.OrganisationID,
			}).Count(&totSpaces)

			if totSpaces >= permission.Spaces && permission.Spaces > 0 {
				errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot create more spaces", http.StatusUnprocessableEntity)))
				return
			}
		}
	}

	var spaceSlug string
	if space.Slug != "" && slug.Check(space.Slug) {
		spaceSlug = space.Slug
	} else {
		spaceSlug = slug.Make(space.Name)
	}

	result := model.Space{
		Name:              space.Name,
		SiteTitle:         space.SiteTitle,
		Slug:              slug.ApproveSpaceSlug(spaceSlug),
		Description:       space.Description,
		TagLine:           space.TagLine,
		SiteAddress:       space.SiteAddress,
		VerificationCodes: space.VerificationCodes,
		SocialMediaURLs:   space.SocialMediaURLs,
		OrganisationID:    space.OrganisationID,
		ContactInfo:       space.ContactInfo,
	}

	tx := config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Begin()
	err = tx.Create(&result).Error

	if err != nil {
		tx.Rollback()
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	if viper.GetBool("create_super_organisation") {
		// Create SpacePermission for super organisation
		var spacePermission model.SpacePermission
		if superOrgID == space.OrganisationID {
			spacePermission = model.SpacePermission{
				SpaceID: result.ID,
				Charts:  -1,
			}
		} else {
			spacePermission = model.SpacePermission{
				SpaceID: result.ID,
				Charts:  viper.GetInt64("default_number_of_charts"),
			}
		}
		var spacePermContext config.ContextKey = "space_perm_user"
		if err = tx.WithContext(context.WithValue(r.Context(), spacePermContext, uID)).Create(&spacePermission).Error; err != nil {
			tx.Rollback()
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.DBError()))
			return
		}

	}

	tx.Commit()
	renderx.JSON(w, http.StatusCreated, result)
}
