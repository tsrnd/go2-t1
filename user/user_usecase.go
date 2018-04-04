package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	GetAllUser() ([]ResponseUser, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db         *gorm.DB
	repository RepositoryInterface
}

// GetAllUser comment
func (u *Usecase) GetAllUser() ([]ResponseUser, error) {
	user, err := u.repository.GetAllUser()
	if err != nil {
		err = utils.ErrorsWrap(err, "Error")
	}
	responseUser := []ResponseUser{}

	for _, v := range user {
		responseUser = append(responseUser, ResponseUser{ID: v.ID, Username: v.Username, Password: v.Password, Phone: v.Phone})
	}
	return responseUser, err
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}
