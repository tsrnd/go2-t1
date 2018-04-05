package user

// PostRegisterByDeviceRequest struct.
type PostRegisterByDeviceRequest struct {
	DeviceID string `form:"device_id" validate:"required,uuid"`
}

// PostRegisterRequest struct.
type PostRegisterRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Phone    string `form:"phone" validate:"numeric"`
	Avatar   string `form:"avatar"`
}
