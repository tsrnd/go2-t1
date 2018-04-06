package user

// PostRegisterByDeviceRequest struct.
type PostRegisterByDeviceRequest struct {
	DeviceID string `form:"device_id" validate:"required,uuid"`
}

// PostRegisterRequest struct.
type PostRegisterRequest struct {
	Username string `form:"username" validate:"required,min=6,max=13"`
	Password string `form:"password" validate:"required,min=6"`
}
