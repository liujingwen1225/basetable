package v1

import "basetable.com/internal/basetable/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.UserService
)

var ApiGroupApp = new(ApiGroup)
