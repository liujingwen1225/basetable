package user

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	"basetable.com/pkg/api"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/token"
	"context"
	"github.com/jinzhu/copier"
	"regexp"
)

type IUserBiz interface {
	Create(cxt context.Context, request *v1.CreateUserRequest) error
	Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error)
	List(cxt context.Context, request *api.PageRequest) ([]*model.UserM, error)
	Deleted(cxt context.Context, userIds int) error
	GetOne(cxt context.Context, id int) (*model.UserM, error)
	Update(cxt context.Context, request *v1.UpdateUserRequest) error
}

var (
	_     IUserBiz = &User{}
	query          = store.Q
)

type User struct {
}

func NewUserBiz() *User {
	return &User{}
}

func (u *User) Create(cxt context.Context, request *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	if err := query.UserM.Create(&userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (u *User) Login(cxt context.Context, request *v1.UserLoginRequest) (*v1.LoginResponse, error) {
	user, err := query.UserM.Where(query.UserM.Username.Eq(request.Username)).First()
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

func (u *User) List(cxt context.Context, page *api.PageRequest) ([]*model.UserM, error) {
	offset := (page.Page - 1) * page.PageSize
	users, _, err := query.UserM.FindByPage(offset, page.PageSize)
	return users, err
}

func (u *User) Deleted(cxt context.Context, userId int) error {
	user := model.UserM{BaseModelM: model.BaseModelM{ID: userId}}
	_, err := query.UserM.Delete(&user)
	return err
}

func (u *User) GetOne(cxt context.Context, id int) (*model.UserM, error) {
	user, err := query.UserM.Where(query.UserM.ID.Eq(id)).First()
	return user, err
}

func (u *User) Update(cxt context.Context, user *v1.UpdateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, user)
	if _, err := query.UserM.Updates(&userM); err != nil {
		return errno.ErrUserCreate
	}
	return nil
}
