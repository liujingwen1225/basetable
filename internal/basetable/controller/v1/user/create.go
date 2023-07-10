package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	v1 "basetable.com/pkg/api/basetable/v1"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var r v1.CreateUserRequest
	// 绑定参数
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	// 参数校验
	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	// 创建业务
	if err := ctrl.biz.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
