package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/usecase"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	Register(request PostRegisterRequest) error
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

func (u *Usecase) Register(request PostRegisterRequest) error {
	tx := u.db.Begin()
	_, err := u.repository.Create(request.Username, request.Password, request.Phone, request.Avatar, tx)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
