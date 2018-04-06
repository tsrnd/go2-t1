package user

import (
	"net/http"

	"github.com/tsrnd/trainning/infrastructure"
	"github.com/tsrnd/trainning/shared/handler"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/usecase"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	usecase UsecaseInterface
}

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := PostRegisterRequest{}
	err := h.Parse(r, &request)
	if err != nil {
		common := CommonResponse{Message: "Parse request error.", Errors: []string{}}
		h.StatusBadRequest(w, common)
		return
	}
	if err = h.Validate(w, request); err != nil {
		return
	}
	response, err := h.usecase.Register(request)
	if err != nil {
		common := CommonResponse{Message: "Internal server error response.", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// user set.
	userRepo := NewRepository(br, s.Master, s.Read, c.Conn)
	userUsecase := NewUsecase(bu, s.Master, userRepo)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: userUsecase}
}
