package model

// KetoPolicy model
type KetoPolicy struct {
	ID          string   `json:"id"`
	Subjects    []string `json:"subjects"`
	Actions     []string `json:"actions"`
	Resources   []string `json:"resources"`
	Effect      string   `json:"effect"`
	Description string   `json:"description"`
}

// Permission model
type Permission struct {
	Resource string   `json:"resource"`
	Actions  []string `json:"actions"`
}

// Policy model
type Policy struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Permissions []Permission  `json:"permissions"`
	Subjects    []interface{} `json:"subjects"`
}
