package collections

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/pkg/api"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) List(c *gin.Context) {
	var pageRequest api.PageRequest
	if err := c.ShouldBindJSON(&pageRequest); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}
	list, err := ctrl.biz.Collections().List(c, &pageRequest)
	core.WriteResponse(c, err, list)
}
