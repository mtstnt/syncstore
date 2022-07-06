package files

import (
	"syncstore/domains/user"

	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Name   string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	UserID uint      `gorm:"column:user_id"`
	User   user.User `json:"user"`
}
