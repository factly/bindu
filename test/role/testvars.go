package role

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

var Data = map[string]interface{}{
	"name":        "admin",
	"description": "",
	"users":       []string{"1"},
}

var basePath = "/roles"
var path = "/roles/{role_id}"
