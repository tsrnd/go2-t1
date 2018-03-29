package Handlers

import (
	. "go-t1/app/Models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetListUser() []User
	DeleteUser(id uint32) bool
}

type userRepository struct {
	DB *gorm.DB
}

func (u userRepository) GetListUser() []User {
	users := []User{}
	u.DB.Find(&users)
	return users
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) DeleteUser(id uint32) bool {
	u.DB.Where("id = ?", id).Delete(&User{})
	return true
}
