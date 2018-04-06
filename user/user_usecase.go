package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	Register(request PostRegisterRequest) (CommonResponse, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

func (u *Usecase) Register(request PostRegisterRequest) (response CommonResponse, err error) {
	response = CommonResponse{}
	tx := u.db.Begin()
	_, err = u.repository.Create(request.Username, request.Password)
	if err != nil {
		tx.Rollback()
	}
	response.Message = "register success"
	tx.Commit()
	return response, utils.ErrorsWrap(err, "repository.Create() error")
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
