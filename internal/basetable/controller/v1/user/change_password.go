package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	v1 "basetable.com/pkg/api/basetable/v1"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) ChangePassword(c *gin.Context) {
	log.C(c).Infow("user password change")
	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

}
