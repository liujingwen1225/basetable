package biz

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/token"
	"context"
	"github.com/jinzhu/copier"
)

type UserBiz interface {
	Create(cxt context.Context, request *v1.UserRequest) (*v1.UserResponse, error)
	Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error)
	Deleted(cxt context.Context, ids []int) error
	GetById(cxt context.Context, id int) (*v1.UserResponse, error)
	List(cxt context.Context, request *v1.ListUserRequest) ([]*v1.UserResponse, error)
	Update(cxt context.Context, request *v1.UpdateUserRequest) (*v1.UserResponse, error)
}

var _ UserBiz = &userBiz{}

type userBiz struct {
	ds store.IStore
}

func (u *userBiz) Update(cxt context.Context, request *v1.UpdateUserRequest) (*v1.UserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	update, err := u.ds.Users().Update(cxt, &userM)
	var res v1.UserResponse
	_ = copier.Copy(&res, update)
	return &res, err
}

func (u *userBiz) GetById(cxt context.Context, id int) (*v1.UserResponse, error) {
	var res v1.UserResponse
	byId, err := u.ds.Users().GetById(cxt, id)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	_ = copier.Copy(&res, byId)
	return &res, err
}

func (u *userBiz) List(cxt context.Context, request *v1.ListUserRequest) ([]*v1.UserResponse, error) {
	var userM model.UserM
	var pagination v1.Pagination
	_ = copier.Copy(&userM, request)
	_ = copier.Copy(&pagination, request)

	list, err := u.ds.Users().List(cxt, &userM, pagination)
	if err != nil {
		return nil, err
	}
	var res []*v1.UserResponse
	_ = copier.Copy(&res, list)
	return res, nil
}

func (u *userBiz) Create(cxt context.Context, request *v1.UserRequest) (*v1.UserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	name, _ := u.ds.Users().GetByUserName(cxt, userM.Username)
	if name != nil {
		return nil, errno.ErrUserAlreadyExist
	}
	if err := u.ds.Users().Create(cxt, &userM); err != nil {
		return nil, err
	}
	var res v1.UserResponse
	_ = copier.Copy(&res, userM)
	return &res, nil
}

func (u *userBiz) Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error) {
	user, err := u.ds.Users().GetByUserName(cxt, request.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
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

func (u *userBiz) Deleted(cxt context.Context, ids []int) error {
	err := u.ds.Users().Delete(cxt, ids)
	return err
}

func NewUserBiz(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}
