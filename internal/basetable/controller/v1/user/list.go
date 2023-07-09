package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/pkg/api"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) List(c *gin.Context) {
	var pageRequest api.PageRequest
	_ = c.ShouldBindJSON(&pageRequest)
	list, err := ctrl.biz.Users().List(c, &pageRequest)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, list)
}
