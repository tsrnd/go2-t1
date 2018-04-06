package user

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/utils"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	Create(string, string) (uint64, error)
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

// CreateUser create user
func (r *Repository) Create(username string, password string) (uint64, error) {
	user := User{Username: username, Password: password}
	err := r.masterDB.FirstOrCreate(&user, User{Username: username}).Error
	if err != nil {
		return user.ID, utils.ErrorsWrap(err, "can't create user")
	}
	return user.ID, nil
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}
