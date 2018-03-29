package Handlers

import (
	"go-t1/app/Models"
	. "go-t1/app/Models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetListUser() []User
	DeleteUser(id uint32) bool
	InsertUser(name string, city string, identityID int64, gender bool)
}

type userRepository struct {
	DB *gorm.DB
}

func (u userRepository) GetListUser() []User {
	users := []User{}
	u.DB.Find(&users)
	return users
}

func (u userRepository) InsertUser(name string, city string, identityID int64, gender bool) {
	users := Models.User{Name: name, City: city, IdentityId: identityID, Gender: gender}
	u.DB.Create(&users)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) DeleteUser(id uint32) bool {
	u.DB.Where("id = ?", id).Delete(&User{})
	return true
}
