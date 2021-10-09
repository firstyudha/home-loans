package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func StaffSeeder(db *gorm.DB) {

	//staff account
	var user User
	user.Username = "staff"
	user.LoginAs = 2

	//hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.MinCost)

	if err != nil {
		log.Fatal(err.Error())
	}

	user.Password = string(passwordHash)

	//check if staff already exist
	err = db.Where("username = ?", user.Username).Find(&user).Error
	if err != nil {
		log.Fatal(err.Error())
	}

	if user.ID > 0 {
		return
	}

	//create if doesn't exist
	err = db.Create(&user).Error

	if err != nil {
		log.Fatal(err.Error())
	}
}
