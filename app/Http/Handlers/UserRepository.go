package Handlers

import (
	. "go-t1/app/Models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetListUser() []User
}

type userRepository struct {
	DB *gorm.DB
}

func (u *userRepository) GetListUser() []User {
	defer u.DB.Close()
	users := []User{}
	u.DB.Find(&users)
	return users
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
