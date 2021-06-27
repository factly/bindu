package policy

var headers = map[string]string{
	"X-Space": "1",
	"X-User":  "1",
}

// valid policy
var policy_test = map[string]interface{}{
	"name": "test policy",
	"permissions": []map[string]interface{}{
		{
			"resource": "policies",
			"actions":  []string{"get", "create", "update", "delete"},
		},
	},
	"users": []string{
		"test@test.com",
	},
}

var undecodable_policy = map[string]interface{}{
	"name":        "test policy",
	"permissions": "none",
	"users": []string{
		"test@test.com",
	},
}

var basePath = "/policies"
var defaultsPath = "/policies/default"
var path = "/policies/{policy_id}"
