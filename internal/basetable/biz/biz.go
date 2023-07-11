package biz

import (
	"basetable.com/internal/basetable/biz/collections"
	"basetable.com/internal/basetable/biz/user"
)

type IBiz interface {
	Users() user.IUserBiz
	Collections() collections.ICollectionsBiz
}

var _ IBiz = &Biz{}

type Biz struct {
}

func New() *Biz {
	return &Biz{}
}

func (b *Biz) Users() user.IUserBiz {
	return user.NewUserBiz()
}

func (b *Biz) Collections() collections.ICollectionsBiz {
	return collections.NewCollectionsBiz()
}
