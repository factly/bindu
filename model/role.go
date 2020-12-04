package model

// KetoRole model
type KetoRole struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Members     []string `json:"members"`
}

// Role model
type Role struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Users       []User `json:"users"`
}
