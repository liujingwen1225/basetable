package collections

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ctrl *Controller) Delete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		core.WriteResponse(c, errno.ErrCollectionsTypeNotFound, nil)
	}
	if err := ctrl.biz.Collections().Delete(c, id); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
