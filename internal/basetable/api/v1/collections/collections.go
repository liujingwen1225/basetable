package collections

//
//import (
//	"basetable.com/internal/basetable/biz"
//	"basetable.com/internal/basetable/store"
//	"basetable.com/internal/pkg/core"
//	"basetable.com/internal/pkg/errno"
//	"basetable.com/internal/pkg/log"
//	v1 "basetable.com/pkg/api/basetable/v1"
//	"github.com/asaskevich/govalidator"
//	"github.com/gin-gonic/gin"
//	"strconv"
//)
//
//type CollectionsController struct {
//	biz biz.IBiz
//}
//
//func New(ds store.IStore) *CollectionsController {
//	return &CollectionsController{biz: biz.NewBiz(ds)}
//}
//
//func (ctrl *CollectionsController) Create(c *gin.Context) {
//	log.C(c).Infow("Create user function called")
//	var r v1.UserRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//	// Validate
//	if _, err := govalidator.ValidateStruct(r); err != nil {
//		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
//		return
//	}
//
//	create, err := ctrl.biz.Users().Create(c, &r)
//	if err != nil {
//		core.WriteResponse(c, err, nil)
//		return
//	}
//
//	core.WriteResponse(c, nil, create)
//}
//
//func (ctrl *CollectionsController) ChangePassword(c *gin.Context) {
//	log.C(c).Infow("user password change")
//	var r v1.ChangePasswordRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//
//}
//
//func (ctrl *CollectionsController) Login(c *gin.Context) {
//	log.C(c).Infow("user login")
//	var r v1.UserLoginRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//	resp, err := ctrl.biz.Users().Login(c, &r)
//	if err != nil {
//		core.WriteResponse(c, err, nil)
//		return
//	}
//
//	core.WriteResponse(c, nil, resp)
//}
//
//func (ctrl *CollectionsController) Update(c *gin.Context) {
//	var r v1.UpdateUserRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//
//	update, err := ctrl.biz.Users().Update(c, &r)
//	if err != nil {
//		core.WriteResponse(c, err, nil)
//		return
//	}
//	core.WriteResponse(c, nil, update)
//}
//
//func (ctrl *CollectionsController) Deleted(c *gin.Context) {
//	var r v1.DeletedUserRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//	if err := ctrl.biz.Users().Deleted(c, r.Ids); err != nil {
//		core.WriteResponse(c, errno.ErrUserDeleted, nil)
//		return
//	}
//	core.WriteResponse(c, nil, nil)
//}
//
//func (ctrl *CollectionsController) GetById(c *gin.Context) {
//	strId := c.Param("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		core.WriteResponse(c, errno.ErrUserIdType, nil)
//		return
//	}
//	byId, err := ctrl.biz.Users().GetById(c, id)
//	if err != nil {
//		core.WriteResponse(c, errno.ErrUserNotFound, nil)
//		return
//	}
//	core.WriteResponse(c, nil, byId)
//}
//
//func (ctrl *CollectionsController) List(c *gin.Context) {
//	var r v1.ListUserRequest
//	if err := c.ShouldBindJSON(&r); err != nil {
//		core.WriteResponse(c, errno.ErrBind, nil)
//		return
//	}
//	list, err := ctrl.biz.Users().List(c, &r)
//	if err != nil {
//		core.WriteResponse(c, errno.ErrUserDeleted, nil)
//		return
//	}
//	core.WriteResponse(c, nil, list)
//}
