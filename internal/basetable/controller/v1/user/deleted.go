package user

import (
	"basetable.com/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Deleted(c *gin.Context) {
	id := c.GetInt("id")
	err := ctrl.b.Users().Deleted(c, id)
	core.WriteResponse(c, err, nil)
}
