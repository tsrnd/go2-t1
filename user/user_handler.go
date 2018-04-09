package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
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

// RegisterByDevice to register user ID which originates from Device ID.
//
// "First": Search User from Entity by Device ID.
// "Second": If User record exists,move to step "Finally".
// "Third": If User record does not exist, register device ID to Entity.
// "Finally":store User_ID acquired from Entity to JSON Web Token (JWT).
func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := LoginRequest{}
	err := h.ParseMultipart(r, &request)
	if err != nil {
		common := CommonResponse{Message: "Parse request error.", Errors: nil}
		h.StatusBadRequest(w, common)
		return
	}

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	// request login by uuid.
	response, err := h.usecase.Login(request)
	if err != nil {
		common := CommonResponse{Message: "Internal server error response.", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := RegisterRequest{}
	err := h.ParseMultipart(r, &request)
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

func (h *HTTPHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		common := CommonResponse{Message: "Parse Request Error.", Errors: []string{}}
		h.StatusBadRequest(w, common)
	}

	err = h.usecase.Destroy(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			common := CommonResponse{Message: "User Not Found.", Errors: []string{}}
			h.StatusNotFoundRequest(w, common)
		}
		common := CommonResponse{Message: "Internal Server Error.", Errors: []string{}}
		h.StatusServerError(w, common)
	}

	common := CommonResponse{Message: "Delete Success.", Errors: []string{}}
	h.ResponseJSON(w, common)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// user set.
	userRepo := NewRepository(br, s.Master, s.Read, c.Conn)
	userUsecase := NewUsecase(bu, s.Master, userRepo)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: userUsecase}
}
