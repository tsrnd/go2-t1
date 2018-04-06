package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/authentication"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	Login(LoginRequest) (LoginReponse, error)
	Register(RegisterRequest) (CommonResponse, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

func (u *Usecase) Login(request LoginRequest) (response LoginReponse, err error) {
	// var userID int64
	response = LoginReponse{}
	user, err := u.repository.Find(request.Username, request.Password)
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

func (u *Usecase) Register(request RegisterRequest) (response CommonResponse, err error) {
	response = CommonResponse{}

	if request.Password != request.RepeatPassword {
		return response, utils.ErrorsWrap(err, "password not match")
	}

	response.Message = "Register success"
	user := User{
		Username: request.Username,
		Password: request.Password,
	}
	_, err = u.repository.Create(user)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.Create() error")
	}
	return
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
