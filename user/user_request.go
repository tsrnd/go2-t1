package user

import "mime/multipart"

// PostRegisterByDeviceRequest struct.
type UserLoginRequest struct {
	Username string `form:"username" validate:"required,min=3"`
	Password string `form:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Username       string `form:"username" validate:"required,min=3"`
	Password       string `form:"password" validate:"required"`
	RepeatPassword string `form:"repeat_password" validate:"required"`
}

type Image struct {
	ImageFile multipart.File
	FileType  string `form:"image type" validate:"omitempty,eq=image/bmp|eq=image/dib|eq=image/jpeg|eq=image/jp2|eq=image/png|eq=image/webp|eq=image/x-portable-anymap|eq=image/x-portable-bitmap|eq=image/x-portable-graymap|eq=image/x-portable-pixmap|eq=image/x-cmu-raster|eq=image/tiff|eq=image/gif"`
	FileSize  int64  `form:"image size" validate:"omitempty,gt=0,max=10485760"`
	FileName  string
}
