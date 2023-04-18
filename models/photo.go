package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo Model
type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Your title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Photo url is required"`
	UserID   uint
	User     *User
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

// hook before create photo
func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	err = nil
	return
}

// hook before update photo
func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
