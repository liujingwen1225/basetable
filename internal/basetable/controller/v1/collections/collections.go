package collections

import (
	"basetable.com/internal/basetable/biz"
)

type Controller struct {
	biz biz.IBiz
}

func New() *Controller {
	return &Controller{biz: biz.New()}
}
