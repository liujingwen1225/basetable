package store

import (
	"basetable.com/internal/pkg/model"
	"basetable.com/pkg/api"
	"basetable.com/pkg/db"
	"context"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(cxt context.Context, user *model.UserM) error
	GetOne(cxt context.Context, user *model.UserM) (*model.UserM, error)
	GetPage(cxt context.Context, request *api.PageRequest) (*[]model.UserM, error)
	Update(cxt context.Context, user *model.UserM) error
	Deleted(cxt context.Context, userIds int) error
}

var _ UserStore = &users{}

type users struct {
	db *gorm.DB
}

func newUserStore(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(cxt context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func (u *users) GetOne(cxt context.Context, userInput *model.UserM) (*model.UserM, error) {
	var userOutput model.UserM
	if err := u.db.Where(&userInput).First(&userOutput).Error; err != nil {
		return nil, err
	}
	return &userOutput, nil
}

func (u *users) Update(cxt context.Context, user *model.UserM) error {
	return u.db.Save(&user).Error
}

func (u *users) Deleted(cxt context.Context, userId int) error {
	return u.db.Where("id = ?", userId).Delete(&model.UserM{}).Error
}

func (u *users) GetPage(cxt context.Context, pageRequest *api.PageRequest) (*[]model.UserM, error) {
	var userOutput []model.UserM
	if err := u.db.Scopes(db.Paginate(pageRequest)).Find(&userOutput).Error; err != nil {
		return nil, err
	}
	return &userOutput, nil
}
