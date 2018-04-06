package user

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	infra "github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/utils"
)

const (
	// S3Path is dicretory.
	S3Path = "item"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	CheckUserByID(id uint64) error
	AddImageToS3(Image) (string, error)
	UpdateUser(password, phone, avatar string, id uint64) error
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
	// s3
	S3Func infra.NewS3RequestFunc
}

// CheckUserByID check  if user has already exist
func (r *Repository) CheckUserByID(id uint64) error {
	if err := r.readDB.First(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// AddImageToS3 function add image
func (r *Repository) AddImageToS3(image Image) (string, error) {
	bucketName := infra.GetConfigString("objectstorage.bucketname")
	if image.ImageFile == nil {
		return "", utils.ErrorsWrap(errors.New("no file detected"), "can't find uploadfile")
	}
	objectName := utils.GetObjectPath(infra.Storage, S3Path, image.FileName)

	params, err := r.S3Func().SetParam(image.ImageFile, bucketName, objectName, image.FileType, s3.BucketCannedACLPublicReadWrite).UploadToS3()
	r.Logger.Debug(params.String())
	if err != nil {
		return "", utils.ErrorsWrap(err, "can't put S3")
	}
	return utils.GetStorageURL(infra.Storage, infra.Endpoint, infra.Secure, bucketName, objectName, infra.Region), nil
}

// UpdateUser update infomartion of user which is given by id
func (r *Repository) UpdateUser(password, phone, avatar string, id uint64) error {
	var err error

	user := &User{Password: password, Phone: phone, Avatar: avatar}

	updateUser := r.masterDB.Model(&User{}).Where("id = ?", id).Update(user)

	if updateUser.Error != nil {
		err = utils.ErrorsWrap(updateUser.Error, "Can't create error.")
	}
	return err
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *gorm.DB, read *gorm.DB, redis *redis.Conn, s3RequestFunc infra.NewS3RequestFunc) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read, redis: redis, S3Func: s3RequestFunc}
}
