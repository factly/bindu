package test

import "time"

// Dummy response body for the mock server requesting organisation data
// Endpoint this is sent for is /organisations
var Dummy_Org = map[string]interface{}{
	"id":         1,
	"created_at": time.Now(),
	"updated_at": time.Now(),
	"deleted_at": nil,
	"title":      "test org",
	"slug":       "test-org",
	"permission": map[string]interface{}{
		"id":              1,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
		"deleted_at":      nil,
		"user_id":         1,
		"user":            nil,
		"organisation_id": 1,
		"organisation":    nil,
		"role":            "owner",
	},
}

var PaiganatedOrg = map[string]interface{}{
	"nodes": []interface{}{
		Dummy_Org,
	},
	"total": 1,
}

var Dummy_Org_Member_List = []map[string]interface{}{
	map[string]interface{}{
		"id":         1,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"title":      "test org",
		"permission": map[string]interface{}{
			"id":              1,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
			"deleted_at":      nil,
			"user_id":         1,
			"user":            nil,
			"organisation_id": 1,
			"organisation":    nil,
			"role":            "member",
		},
	},
}

var Dummy_OrgList = []map[string]interface{}{
	Dummy_Org,
}

// Dummy response for the mock server requesting list of authors
// Endpoint this is sent for is /organisations/[id]/users
var Dummy_AuthorList = []map[string]interface{}{
	{
		"id":         1,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"email":      "abc@abc.com",
		"kid":        "",
		"first_name": "abc",
		"last_name":  "cba",
		"birth_date": time.Now(),
		"gender":     "male",
		"permission": map[string]interface{}{
			"id":              1,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
			"deleted_at":      nil,
			"user_id":         1,
			"user":            nil,
			"organisation_id": 1,
			"organisation":    nil,
			"role":            "owner",
		},
	},
	{
		"id":         2,
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"deleted_at": nil,
		"email":      "def@def.com",
		"kid":        "",
		"first_name": "def",
		"last_name":  "fed",
		"birth_date": time.Now(),
		"gender":     "male",
		"permission": map[string]interface{}{
			"id":              2,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
			"deleted_at":      nil,
			"user_id":         2,
			"user":            nil,
			"organisation_id": 1,
			"organisation":    nil,
			"role":            "member",
		},
	},
}

var Dummy_KetoPolicy = []map[string]interface{}{
	{
		"id":          "id:org:1:app:bindu:space:1:test-policy-4",
		"description": "",
		"subjects": []string{
			"1",
			"2",
		},
		"resources": []string{
			"resources:org:1:app:bindu:space:1:categories",
			"resources:org:1:app:bindu:space:1:tags",
		},
		"actions": []string{
			"actions:org:1:app:bindu:space:1:categories:get",
			"actions:org:1:app:bindu:space:1:categories:create",
			"actions:org:1:app:bindu:space:1:tags:update",
			"actions:org:1:app:bindu:space:1:tags:delete",
		},
		"effect":     "allow",
		"conditions": nil,
	},
	{
		"id":          "id:org:1:app:bindu:space:1:test-policy-0",
		"description": "",
		"subjects": []string{
			"1",
		},
		"resources": []string{
			"resources:org:12:app:bindu:space:18:policies",
		},
		"actions": []string{
			"actions:org:12:app:bindu:space:18:policies:get",
			"actions:org:12:app:bindu:space:18:policies:create",
			"actions:org:12:app:bindu:space:18:policies:update",
			"actions:org:12:app:bindu:space:18:policies:delete",
		},
		"effect":     "allow",
		"conditions": nil,
	},
}

var Dummy_Role = map[string]interface{}{
	"id": "roles:org:1:admin",
	"members": []string{
		"1",
	},
}

// Dummy single policy
var Dummy_SingleMock = map[string]interface{}{
	"id":          "id:org:1:app:bindu:space:1:test-policy-0",
	"description": "",
	"subjects": []string{
		"1",
	},
	"resources": []string{
		"resources:org:12:app:bindu:space:18:policies",
	},
	"actions": []string{
		"actions:org:12:app:bindu:space:18:policies:get",
		"actions:org:12:app:bindu:space:18:policies:create",
		"actions:org:12:app:bindu:space:18:policies:update",
		"actions:org:12:app:bindu:space:18:policies:delete",
	},
	"effect":     "allow",
	"conditions": nil,
}

var ReturnUpdate = map[string]interface{}{
	"updateId": 1,
}

var MeiliHits = map[string]interface{}{
	"hits": []map[string]interface{}{
		{
			"object_id":   "post_3",
			"kind":        "post",
			"description": "This is a test post with claim",
			"id":          1,
			"slug":        "test-post-2",
			"space_id":    1,
			"title":       "Test Post",
			"category_ids": []uint{
				2,
			},
			"excerpt":        "Test post with claim",
			"is_featured":    true,
			"is_highlighted": true,
			"is_sticky":      true,
			"published_date": -62135596800,
			"status":         "draft",
			"subtitle":       "Test Post",
			"tag_ids": []uint{
				42,
			},
			"claim_ids": []uint{
				5,
			},
			"format_id": 3,
		},
		{
			"object_id":       "claim_2",
			"kind":            "claim",
			"description":     "This is a test claim",
			"id":              2,
			"slug":            "test-claim",
			"space_id":        1,
			"title":           "Test Claim",
			"checked_date":    1598959138,
			"claim_date":      -62135596800,
			"claim_sources":   "secret sources",
			"claimant_id":     2,
			"rating_id":       2,
			"review":          "Bad review",
			"review_sources":  "Good sources",
			"review_tag_line": "Bad review good sources",
		},
	},
	"offset":           0,
	"limit":            20,
	"nbHits":           7,
	"exhaustiveNbHits": false,
	"processingTimeMs": 2,
	"query":            "test",
}

var EmptyMeili = map[string]interface{}{
	"hits":             []map[string]interface{}{},
	"offset":           0,
	"limit":            20,
	"nbHits":           0,
	"exhaustiveNbHits": false,
	"processingTimeMs": 2,
	"query":            "test",
}
