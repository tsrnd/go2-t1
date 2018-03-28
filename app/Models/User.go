package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Id         uint32 `sql:"type:integer;primary key"`
	Name       string `sql:"type:text"`
	City       string `sql:"type:text"`
	IdentityId int64  `sql:"type:integer"`
	Gender     bool   `sql:"type:boolean"`
}

func (User) TableName() string {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "golang." + defaultTableName
	}
	return "users"
}
