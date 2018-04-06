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
	_, err = u.repository.Create(request.Username, request.Password)
	response.Message = "register success"
	return response, utils.ErrorsWrap(err, "repository.Create() error")
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
