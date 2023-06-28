package user

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/token"
	"context"
	"github.com/jinzhu/copier"
	"regexp"
)

type UserBiz interface {
	Create(cxt context.Context, request *v1.CreateUserRequest) error
	Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error)
}

var _ UserBiz = &userBiz{}

type userBiz struct {
	ds store.IStore
}

func NewUserBiz(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (u *userBiz) Create(cxt context.Context, request *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	if err := u.ds.Users().Create(cxt, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (u *userBiz) Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error) {
	userInput := &model.UserM{Username: request.Username}
	user, err := u.ds.Users().Get(cxt, userInput)
	if err != nil {
		return nil, err
	}
	if user.Password != request.Password {
		return nil, errno.ErrPasswordIncorrect
	}
	sign, err := token.Sign(user.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}
	return &v1.LoginResponse{Token: sign}, nil
}
