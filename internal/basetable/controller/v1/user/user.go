package user

import (
	"basetable.com/internal/basetable/biz"
	"basetable.com/internal/basetable/store"
)

type Controller struct {
	b biz.IBiz
}

func New(ds store.IStore) *Controller {
	return &Controller{b: biz.NewBiz(ds)}
}
