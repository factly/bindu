package theme

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/renderx"
	"github.com/factly/x/validationx"
)

// create - Create theme
// @Summary Create theme
// @Description Create theme
// @Tags Theme
// @ID add-theme
// @Consume json
// @Produce json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param Theme body theme true "Theme Object"
// @Success 201 {object} model.Theme
// @Failure 400 {array} string
// @Router /themes [post]
func create(w http.ResponseWriter, r *http.Request) {

	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	uID, err := middlewarex.GetUser(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	theme := &theme{}

	err = json.NewDecoder(r.Body).Decode(&theme)
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DecodeError()))
		return
	}

	validationError := validationx.Check(theme)

	if validationError != nil {
		loggerx.Error(errors.New("validation error"))
		errorx.Render(w, validationError)
		return
	}

	// Store HTML description
	var description string
	if len(theme.Description.RawMessage) > 0 && !reflect.DeepEqual(theme.Description, util.NilJsonb()) {
		description, err = util.HTMLDescription(theme.Description)
		if err != nil {
			loggerx.Error(err)
			errorx.Render(w, errorx.Parser(errorx.GetMessage("cannot parse theme description", http.StatusUnprocessableEntity)))
			return
		}
	}

	result := &model.Theme{
		Name:            theme.Name,
		Config:          theme.Config,
		Description:     theme.Description,
		HtmlDescription: description,
		SpaceID:         uint(sID),
	}

	err = config.DB.WithContext(context.WithValue(r.Context(), userContext, uID)).Model(&model.Theme{}).Create(&result).Error

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusCreated, result)
}
