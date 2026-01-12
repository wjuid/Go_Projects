package models

import "gorm.io/gorm"

// examples/models/user.go
type User struct {
	gorm.Model
	Name string
	Age  int
	Pets []Pet `gorm:"many2many:user_pets"`
}

type Pet struct {
	gorm.Model
	Name string
}
