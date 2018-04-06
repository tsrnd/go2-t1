package user

import "mime/multipart"

// UpdateUserRequest struct.
type UpdateUserRequest struct {
	Password string `form:"password" validate:"required"`
	Phone    string `form:"phone" validate:"omitempty,min=10,max=11,numeric"`
	Avatar   Image  `form:"avatar"`
}

// LoginRequest struct.
type LoginRequest struct {
	Username string `form:"username" validate:"required,min=3"`
	Password string `form:"password" validate:"required"`
}

// RegisterRequest struct
type RegisterRequest struct {
	Username       string `form:"username" validate:"required,min=3"`
	Password       string `form:"password" validate:"required"`
	RepeatPassword string `form:"repeat_password" validate:"required"`
}

// Image struct for image information and storage location
type Image struct {
	ImageFile multipart.File
	FileType  string `form:"image type" validate:"omitempty,eq=image/bmp|eq=image/dib|eq=image/jpeg|eq=image/jp2|eq=image/png|eq=image/webp|eq=image/x-portable-anymap|eq=image/x-portable-bitmap|eq=image/x-portable-graymap|eq=image/x-portable-pixmap|eq=image/x-cmu-raster|eq=image/tiff|eq=image/gif"`
	FileSize  int64  `form:"image size" validate:"omitempty,gt=0,max=10485760"`
	FileName  string
}

// PostRegisterRequest struct.
type PostRegisterRequest struct {
	Username string `form:"username" validate:"required,min=6,max=13"`
	Password string `form:"password" validate:"required,min=6"`
}
