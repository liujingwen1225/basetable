package biz

import (
	"basetable.com/internal/basetable/biz/user"
)

type IBiz interface {
	Users() user.IUserBiz
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
