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
	UpdateUser(req UpdateUserRequest, id uint64) (string, error)
}

// UpdateUser update
func (u *Usecase) UpdateUser(req UpdateUserRequest, id uint64) (string, error) {
	err := u.repository.CheckUserByID(id)
	if err == gorm.ErrRecordNotFound {
		return "User is not Exist", err
	}
	avatarURL, err := u.repository.AddImageToS3(req.Avatar)
	if err != nil {
		return "Upload Image to S3 fail", err
	}
	err = u.repository.UpdateUser(req.Password, req.Phone, avatarURL, id)
	if err != nil {
		err = utils.ErrorsWrap(err, "Error")
		return "repository.UpdateUser error", err
	}
	return "Update Information of User successfully", err

}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

// Login to user access into app
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

// Register usercase
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
