package chart

import (
	"errors"
	"time"

	"github.com/factly/bindu-server/model"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// chart request model
type chart struct {
	Title            string         `json:"title" validate:"required,min=3,max=50"`
	Subtitle         string         `json:"subtitle"`
	Slug             string         `json:"slug"`
	URL              string         `json:"url"`
	Description      postgres.Jsonb `json:"description"`
	Status           string         `json:"status"`
	FeaturedMediumID uint           `json:"featured_medium_id"`
	TemplateID       uint           `json:"template_id" validate:"required"`
	ThemeID          uint           `json:"theme_id" validate:"required"`
	PublishedDate    time.Time      `json:"published_date"`
	OrganisationID   uint           `json:"organisation_id"`
}

// Router - Group of chart router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{chart_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}

// CheckOrganisation - validation for medium, theme & templates
func (c *chart) CheckOrganisation(tx *gorm.DB) (e error) {

	if c.FeaturedMediumID > 0 {
		medium := model.Medium{}
		medium.ID = c.FeaturedMediumID

		err := tx.Model(&model.Medium{}).Where(model.Medium{
			OrganisationID: c.OrganisationID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("medium do not belong to same space")
		}
	}

	if c.ThemeID > 0 {
		theme := model.Theme{}
		theme.ID = c.ThemeID

		err := tx.Model(&model.Theme{}).Where(model.Theme{
			OrganisationID: c.OrganisationID,
		}).First(&theme).Error

		if err != nil {
			return errors.New("theme do not belong to same space")
		}
	}

	if c.TemplateID > 0 {
		template := model.Template{}
		template.ID = c.TemplateID

		err := tx.Model(&model.Template{}).Where(model.Template{
			OrganisationID: c.OrganisationID,
		}).First(&template).Error

		if err != nil {
			return errors.New("template do not belong to same space")
		}
	}

	return nil
}
