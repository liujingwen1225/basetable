package user

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ctrl *Controller) Deleted(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		core.WriteResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	erro := ctrl.b.Users().Deleted(c, id)
	core.WriteResponse(c, erro, nil)
}
