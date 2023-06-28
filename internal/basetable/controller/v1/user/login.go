package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	v1 "basetable.com/pkg/api/basetable/v1"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Login(c *gin.Context) {
	log.C(c).Infow("user login")
	var r v1.UserLoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}
