package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Age   int
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}

func (User) TableName() string {
	return "users"
}
