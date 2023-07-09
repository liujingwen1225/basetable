package user

import (
	"basetable.com/internal/pkg/core"
	v1 "basetable.com/pkg/api/basetable/v1"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Update(c *gin.Context) {
	var requst v1.UserRequest
	_ = c.ShouldBindJSON(&requst)
	if err := ctrl.biz.Users().Update(c, &requst); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
