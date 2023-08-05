package service

import (
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/db"
	"basetable.com/pkg/token"
	"github.com/jinzhu/copier"
)

type UserService struct {
}

func (u *UserService) Update(request *v1.UpdateUserRequest) (*v1.UserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	if err := db.DB.Updates(&userM).Error; err != nil {
		return nil, err
	}
	var res v1.UserResponse
	_ = copier.Copy(&res, userM)
	return &res, nil
}

func (u *UserService) GetById(id int) (*v1.UserResponse, error) {
	var userM *model.UserM
	if err := db.DB.Where("id = ?", id).First(&userM).Error; err != nil {
		return nil, errno.ErrUserNotFound
	}
	var res v1.UserResponse
	_ = copier.Copy(&res, userM)
	return &res, nil
}

func (u *UserService) List(request *v1.ListUserRequest) ([]*v1.UserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	offset, limit := request.Pagination.GetPage()
	var userList []*model.UserM
	var count int64
	if err := db.DB.Offset(offset).Limit(limit).Find(&userList).Count(&count).Error; err != nil {
		return nil, err
	}
	var res []*v1.UserResponse
	_ = copier.Copy(&res, userList)
	return res, nil
}

func (u *UserService) Create(request *v1.UserRequest) (*v1.UserResponse, error) {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	var Exist int64
	db.DB.Where("username = ? ", userM.Username).Count(&Exist)
	if Exist != 0 {
		return nil, errno.ErrUserAlreadyExist
	}
	if err := db.DB.Create(&userM).Error; err != nil {
		return nil, err
	}
	var res v1.UserResponse
	_ = copier.Copy(&res, userM)
	return &res, nil
}

func (u *UserService) Login(request *v1.UserLoginRequest) (*v1.LoginResponse, error) {
	var user model.UserM
	if err := db.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
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

func (u *UserService) Deleted(ids []int) error {
	return db.DB.Delete(&model.UserM{}, ids).Error
}
