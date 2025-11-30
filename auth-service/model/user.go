package model

import "gorm.io/gorm"

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"userpassword"`
}

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Userpassword []byte
}
