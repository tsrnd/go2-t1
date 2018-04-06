package user

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/usecase"
	"github.com/tsrnd/trainning/shared/utils"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	usecase UsecaseInterface
}

// UpdateUser handler
func (h *HTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	request := UpdateUserRequest{}
	err := h.ParseMultipart(r, &request)
	if err != nil {
		common := CommonResponse{Message: "Parse request error.", Errors: nil}
		h.StatusBadRequest(w, common)
		return
	}
	f, fh, _ := r.FormFile("avatar")

	if f != nil {
		defer func() {
			err = f.Close()
			if err != nil {
				h.Logger.Error("image_file can't close file.")
			}
		}()

		request.Avatar.ImageFile = f
		fileContentType, err := h.GetFileHeaderContentType(f)
		if err == nil {
			request.Avatar.FileType = fileContentType
			request.Avatar.FileSize = fh.Size
		} else {
			request.Avatar.FileType = "error"
			request.Avatar.FileSize = fh.Size
		}
	}

	//  validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	if request.Avatar.ImageFile != nil {
		// save file.
		filename, err := h.GetRandomFileName("upload_", fh.Filename)
		if err != nil {
			h.Logger.Error("can't get random file name.")
			common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
			h.StatusServerError(w, common)
		}
		request.Avatar.FileName = filename
	}

	ResponseUpdateUser, err := h.usecase.UpdateUser(request, uint64(UserID))
	if err == gorm.ErrRecordNotFound {
		h.StatusBadRequest(w, ResponseUpdateUser)
		return
	}
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecaseInterface.GetOutfitByImName() error")
		common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, ResponseUpdateUser)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL, c *infrastructure.Cache, s3 *infrastructure.S3) *HTTPHandler {
	// user set.
	userRepo := NewRepository(br, s.Master, s.Read, c.Conn, s3.NewRequest)
	userUsecase := NewUsecase(bu, s.Master, userRepo)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: userUsecase}
}
