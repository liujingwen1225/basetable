package biz

import (
	"basetable.com/internal/basetable/biz/user"
	"basetable.com/internal/basetable/store"
)

type IBiz interface {
	Users() user.UserBiz
}

var _ IBiz = &biz{}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Users() user.UserBiz {
	return user.NewUserBiz(b.ds)
}
