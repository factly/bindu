package model

import (
	"errors"

	"github.com/factly/bindu-server/config"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

// Space model
type Space struct {
	config.Base
	Name              string         `gorm:"column:name" json:"name"`
	Slug              string         `gorm:"column:slug" json:"slug"`
	SiteTitle         string         `gorm:"column:site_title" json:"site_title"`
	TagLine           string         `gorm:"column:tag_line" json:"tag_line"`
	Description       string         `gorm:"column:description" json:"description"`
	SiteAddress       string         `gorm:"column:site_address" json:"site_address"`
	LogoID            *uint          `gorm:"column:logo_id;default:NULL" json:"logo_id"`
	Logo              *Medium        `gorm:"foreignKey:logo_id" json:"logo"`
	LogoMobileID      *uint          `gorm:"column:logo_mobile_id;default:NULL" json:"logo_mobile_id"`
	LogoMobile        *Medium        `gorm:"foreignKey:logo_mobile_id" json:"logo_mobile"`
	FavIconID         *uint          `gorm:"column:fav_icon_id;default:NULL" json:"fav_icon_id"`
	FavIcon           *Medium        `gorm:"foreignKey:fav_icon_id" json:"fav_icon"`
	MobileIconID      *uint          `gorm:"column:mobile_icon_id;default:NULL" json:"mobile_icon_id"`
	MobileIcon        *Medium        `gorm:"foreignKey:mobile_icon_id" json:"mobile_icon"`
	VerificationCodes postgres.Jsonb `gorm:"column:verification_codes" json:"verification_codes" swaggertype:"primitive,string"`
	SocialMediaURLs   postgres.Jsonb `gorm:"column:social_media_urls" json:"social_media_urls" swaggertype:"primitive,string"`
	ContactInfo       postgres.Jsonb `gorm:"column:contact_info" json:"contact_info" swaggertype:"primitive,string"`
	OrganisationID    int            `gorm:"column:organisation_id" json:"organisation_id"`
}

var spaceUser config.ContextKey = "space_user"

// BeforeCreate hook
func (space *Space) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userID := ctx.Value(spaceUser)

	if userID == nil {
		return nil
	}
	uID := userID.(int)

	space.CreatedByID = uint(uID)
	space.UpdatedByID = uint(uID)
	return nil
}

// BeforeUpdate checks if all associated mediums are in same space
func (space *Space) BeforeUpdate(tx *gorm.DB) (e error) {
	if space.LogoID != nil && *space.LogoID > 0 {
		medium := Medium{}
		medium.ID = *space.LogoID

		err := tx.Model(&Medium{}).Where(Medium{
			SpaceID: space.ID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("logo do not belong to same space")
		}
	}

	if space.LogoMobileID != nil && *space.LogoMobileID > 0 {
		medium := Medium{}
		medium.ID = *space.LogoMobileID

		err := tx.Model(&Medium{}).Where(Medium{
			SpaceID: space.ID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("logo mobile do not belong to same space")
		}
	}

	if space.FavIconID != nil && *space.FavIconID > 0 {
		medium := Medium{}
		medium.ID = *space.FavIconID

		err := tx.Model(&Medium{}).Where(Medium{
			SpaceID: space.ID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("fav icon do not belong to same space")
		}
	}

	if space.MobileIconID != nil && *space.MobileIconID > 0 {
		medium := Medium{}
		medium.ID = *space.MobileIconID

		err := tx.Model(&Medium{}).Where(Medium{
			SpaceID: space.ID,
		}).First(&medium).Error

		if err != nil {
			return errors.New("mobile icon do not belong to same space")
		}
	}
	return nil
}
