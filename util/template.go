package util

import (
	"encoding/json"
	"html/template"
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// Template template object
var Template *template.Template

// SetupTemplates setups the templates
func SetupTemplates() {
	Template = template.Must(template.New("").Funcs(template.FuncMap{
		"unmar":   Unmarshal,
		"dateFmt": formatDate,
		"dateVal": validateDate,
		"noesc":   noescape,
	}).ParseGlob("web/templates/*"))
}

func Unmarshal(data postgres.Jsonb) map[string]interface{} {
	dataMap := make(map[string]interface{})
	_ = json.Unmarshal(data.RawMessage, &dataMap)
	return dataMap
}

func formatDate(date time.Time) string {
	return date.Format("01/02/2006")
}

func validateDate(date time.Time) bool {
	return !date.Equal(time.Time{})
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}
