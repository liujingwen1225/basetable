package collections

import (
	"basetable.com/internal/pkg/core"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/log"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/util"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) Create(c *gin.Context) {
	log.C(c).Infow("Create collections function called")
	var r v1.CreateCollectionsRequest
	// 绑定参数
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	// 参数校验
	if err := util.Valid(&r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}
	// 创建业务
	if err := ctrl.biz.Collections().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
