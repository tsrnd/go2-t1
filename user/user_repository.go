package user

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/utils"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	Find(string, string) (User, error)
	Create(User) (User, error)
}

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// connect master database.
	masterDB *gorm.DB
	// connect read replica database.
	readDB *gorm.DB
	// redis connect Redis.
	redis *redis.Conn
}

func (r *Repository) Create(user User) (User, error) {
	checkingUser := User{}
	err := r.readDB.Where("username = ?", user.Username).First(&checkingUser)
	if checkingUser.ID > 0 {
		return user, utils.ErrorsWrap(err.Error, "User already exists")
	}
	result := r.masterDB.Create(&user)
	return user, utils.ErrorsWrap(result.Error, "Can't create user")
}

func (r *Repository) Find(username string, password string) (User, error) {
	user := User{Username: username, Password: password}
	err := r.masterDB.First(&user, user).Error
	return user, utils.ErrorsWrap(err, "Can't find this user")
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}
