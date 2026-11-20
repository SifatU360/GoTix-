package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" xml:"name" form:"name" query:"name" validate:"required" gorm:"type:varchar(100); not null"`
	Email    string `json:"email" xml:"email" form:"email" query:"email" validate:"required,email" gorm:"type:varchar(255); uniqueIndex; not null"`
	Password string `json:"password" xml:"password" form:"password" query:"password" validate:"required,min=6" gorm:"type:varchar(100); not null"`
}
