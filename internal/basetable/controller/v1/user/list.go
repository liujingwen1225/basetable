package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/pkg/api"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) List(c *gin.Context) {
	var pageRequest api.PageRequest
	c.ShouldBindJSON(&pageRequest)
	list, err := ctrl.b.Users().List(c, &pageRequest)
	if err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, list)
}
