package model

import (
	"github.com/factly/bindu-server/config"
)

type organisationUser struct {
	config.Base
	Role string `json:"role"`
}

// Organisation model
type Organisation struct {
	config.Base
	Title       string           `json:"title"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Permission  organisationUser `json:"permission"`
}
