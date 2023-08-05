package v1

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	v1 "basetable.com/pkg/api/basetable/v1"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct {
}

func (ctrl *UserApi) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var r v1.UserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	// Validate
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	create, err := userService.Create(&r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, create)
}

func (ctrl *UserApi) ChangePassword(c *gin.Context) {
	log.C(c).Infow("user password change")
	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

}

func (ctrl *UserApi) Login(c *gin.Context) {
	log.C(c).Infow("user login")
	var r v1.UserLoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := userService.Login(&r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *UserApi) Update(c *gin.Context) {
	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	update, err := userService.Update(&r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, update)
}

func (ctrl *UserApi) Deleted(c *gin.Context) {
	var r v1.DeletedUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	if err := userService.Deleted(r.Ids); err != nil {
		core.WriteResponse(c, errno.ErrUserDeleted, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}

func (ctrl *UserApi) GetById(c *gin.Context) {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserIdType, nil)
		return
	}
	byId, err := userService.GetById(id)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	core.WriteResponse(c, nil, byId)
}

func (ctrl *UserApi) List(c *gin.Context) {
	var r v1.ListUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	list, err := userService.List(&r)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserDeleted, nil)
		return
	}
	core.WriteResponse(c, nil, list)
}
