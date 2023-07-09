package user

import (
	"basetable.com/internal/basetable/biz"
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Controller struct {
	biz biz.IBiz
}

func New() *Controller {
	return &Controller{biz: biz.New()}
}

func (ctrl *Controller) GetOne(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	user, err := ctrl.biz.Users().GetOne(c, id)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, user)
}
