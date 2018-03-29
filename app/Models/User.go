package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string `sql:"type:text"`
	City       string `sql:"type:text"`
	IdentityId int64  `sql:"type:integer"`
	Gender     bool   `sql:"type:boolean"`
}

func (User) TableName() string {
	return "golang.users"
}
