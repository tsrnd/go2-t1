package user

import (
	"github.com/tsrnd/trainning/shared/gorm/model"
)

// ----------------------------------------------------------
// Database
// ----------------------------------------------------------

// User table struct.
// http://jinzhu.me/gorm/models.html#model-definition
type User struct {
	ID       uint64 `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username;type:char(36)"`
	Password string `gorm:"column:password;type:char(20)"`
	Phone    string `gorm:"column:phone;type:varchar(11)"`
	Avatar   string `gorm:column:avatar;type:varchar(256)"`
	model.BaseModel
}

// TableName function custom table name.
func (User) TableName() string {
	return "users"
}

// GetCustomClaims get customs claims
func (u User) GetCustomClaims() map[string]interface{} {
	claims := make(map[string]interface{})
	userclaim := struct {
		ID uint64 `json:"id"`
	}{
		ID: u.ID,
	}
	claims["user"] = userclaim
	return claims
}

// GetIdentifier get identifier function
func (u User) GetIdentifier() uint64 {
	return u.ID
}
