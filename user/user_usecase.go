package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/authentication"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	Login(string, string) (PostRegisterByDeviceResponse, error)
	Register(string, string, string) (CommonResponse, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

func (u *Usecase) Login(username string, password string) (response PostRegisterByDeviceResponse, err error) {
	// var userID int64
	response = PostRegisterByDeviceResponse{}
	user, err := u.repository.Find(username, password)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repositoryInterface.Find() error")
	}
	// store user to JWT
	response.Token, err = authentication.GenerateToken(user)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GenerateJWToken() error")
	}
	return
}

func (u *Usecase) Register(username string, password string, repeatPassword string) (response CommonResponse, err error) {
	response = CommonResponse{}

	if password != repeatPassword {
		return response, utils.ErrorsWrap(err, "password not match")
	}

	response.Message = "Register success"
	user := User{
		Username: username,
		Password: password,
	}
	tx := u.db.Begin()
	_, err = u.repository.Create(user)
	tx.Commit()
	if err != nil {
		tx.Rollback()
		return response, utils.ErrorsWrap(err, "repository.Create() error")
	}
	return
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
