package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Username string `gorm:"size:255;unique:true;not null"`
	Password string
	LoginAs  int
}
