package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	UpdateUser(req UpdateUserRequest, id uint64) (string, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
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

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
