package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;type:varchar(255);not null"`
	Password string `json:"-" gorm:"column:username;type:varchar(255);not null"`
	Status   bool   `json:"-" gorm:"column:status;type:boolean default true"`
}
