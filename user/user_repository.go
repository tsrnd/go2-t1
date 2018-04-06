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
	Destroy(id uint64) string
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
	result := r.masterDB.FirstOrCreate(&user, User{Username: user.Username})
	return user, utils.ErrorsWrap(result.Error, "Can't create user")
}

func (r *Repository) Find(username string, password string) (User, error) {
	user := User{Username: username, Password: password}
	err := r.masterDB.First(&user, user).Error
	return user, utils.ErrorsWrap(err, "Can't find this user")
}

func (r *Repository) Destroy(id uint64) string {
	user := User{}
	r.readDB.First(&user, id)
	if user.ID == 0 {
		return "User Not Found"
	}
	err := r.masterDB.Delete(&user)
	if err.Error != nil {
		return "Delete Fail"
	}
	return "Delete Success"
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis}
}
