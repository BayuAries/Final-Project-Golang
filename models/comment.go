package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Model comment
type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint   `json:"photo_id" form:"photo_id"`
	Message string `json:"message" form:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}

// hook before create comment
func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	err = nil
	return
}

// hook before update comment
func (u *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
