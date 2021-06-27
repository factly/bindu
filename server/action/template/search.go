package template

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/factly/bindu-server/action/chart"
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/factly/x/errorx"
	"github.com/factly/x/loggerx"
	"github.com/factly/x/meilisearchx"
	"github.com/factly/x/middlewarex"
	"github.com/factly/x/paginationx"
	"github.com/factly/x/renderx"
)

// list response
type paging struct {
	Total int64            `json:"total"`
	Nodes []model.Template `json:"nodes"`
}

// search - Search all templates
// @Summary Search in all templates
// @Description Search all templates
// @Tags Template
// @ID search-all-templates
// @Produce  json
// @Param X-User header string true "User ID"
// @Param X-Space header string true "Space ID"
// @Param limit query string false "limit per page"
// @Param page query string false "page number"
// @Param category query string false "category"
// @Param q query string false "search query"
// @Success 200 {array} paging
// @Router /templates/search [get]
func search(w http.ResponseWriter, r *http.Request) {
	sID, err := middlewarex.GetSpace(r.Context())
	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.Unauthorized()))
		return
	}

	// Filters
	u, _ := url.Parse(r.URL.String())
	queryMap := u.Query()

	searchQuery := r.URL.Query().Get("q")
	sort := r.URL.Query().Get("sort")

	filters := generateFilters(queryMap["category"])
	filteredTemplateIDs := make([]string, 0)

	if filters != "" {
		filters = fmt.Sprint(filters, " AND space_id=", sID)
	}

	offset, limit := paginationx.Parse(r.URL.Query())

	result := paging{}
	result.Nodes = make([]model.Template, 0)

	if filters != "" || searchQuery != "" {
		// Search posts with filter
		var hits []interface{}
		var res map[string]interface{}

		if searchQuery != "" {
			hits, err = meilisearchx.SearchWithQuery("bindu", searchQuery, filters, "template")
		} else {
			res, err = meilisearchx.SearchWithoutQuery("bindu", filters, "template")
			if _, found := res["hits"]; found {
				hits = res["hits"].([]interface{})
			}
		}
		if err != nil {
			loggerx.Error(err)
			renderx.JSON(w, http.StatusOK, result)
			return
		}

		filteredTemplateIDs = chart.GetIDArray(hits)
		if len(filteredTemplateIDs) == 0 {
			renderx.JSON(w, http.StatusOK, result)
			return
		}
	}

	if sort != "asc" {
		sort = "desc"
	}

	tx := config.DB.Preload("Medium").Preload("Category").Model(&model.Template{}).Order("created_at " + sort)

	if len(filteredTemplateIDs) > 0 {
		err = tx.Where(filteredTemplateIDs).Count(&result.Total).Offset(offset).Limit(limit).Find(&result.Nodes).Error
	} else {
		err = tx.Count(&result.Total).Offset(offset).Limit(limit).Find(&result.Nodes).Error
	}

	if err != nil {
		loggerx.Error(err)
		errorx.Render(w, errorx.Parser(errorx.DBError()))
		return
	}

	renderx.JSON(w, http.StatusOK, result)
}

func generateFilters(categoryIDs []string) string {
	filters := ""
	if len(categoryIDs) > 0 {
		filters = fmt.Sprint(filters, meilisearchx.GenerateFieldFilter(categoryIDs, "category_id"), " AND ")
	}

	if filters != "" && filters[len(filters)-5:] == " AND " {
		filters = filters[:len(filters)-5]
	}

	return filters
}
