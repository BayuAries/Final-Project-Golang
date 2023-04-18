package models

import (
	"MyGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Model for User
type User struct {
	GormModel
	UserName    string       `gorm:"not null;unique" json:"username" form:"username" valid:"required~Your user name is required"`
	Email       string       `gorm:"not null;unique" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password    string       `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         int64        `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(8|150)~Your age minimal is 8"`
	Photos      []Photo      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments    []Comment    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

// hook before create user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
