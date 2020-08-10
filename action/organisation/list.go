package organisation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/errorx"
	"github.com/factly/x/renderx"
)

// list response
type paging struct {
	Total int                  `json:"total"`
	Nodes []model.Organisation `json:"nodes"`
}

// list - Get all organisations
// @Summary Show all organisations
// @Description Get all organisations
// @Tags Organisation
// @ID get-all-organisations
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Organisation header string true "Organisation ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Success 200 {array} model.Organisation
// @Router /organisations [get]
func list(w http.ResponseWriter, r *http.Request) {

	uID, err := util.GetUser(r.Context())

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	req, err := http.NewRequest("GET", os.Getenv("KAVACH_URL")+"/organisations/my", nil)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.InternalServerError()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User", fmt.Sprint(uID))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		errorx.Render(w, errorx.Parser(errorx.NetworkError()))
		return
	}

	defer resp.Body.Close()

	var orgs []interface{}
	json.NewDecoder(resp.Body).Decode(&orgs)

	result := paging{}
	result.Nodes = make([]model.Organisation, 0)

	for _, each := range orgs {
		eachOrg := (each).(map[string]interface{})

		org := model.Organisation{}

		org.ID = int((eachOrg["id"]).(float64))
		org.Title = (eachOrg["title"]).(string)

		result.Nodes = append(result.Nodes, org)
	}

	renderx.JSON(w, http.StatusOK, result)
}
