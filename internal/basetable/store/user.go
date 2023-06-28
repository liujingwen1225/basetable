package store

import (
	"basetable.com/internal/pkg/model"
	"context"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(cxt context.Context, user *model.UserM) error
	Get(cxt context.Context, user *model.UserM) (*model.UserM, error)
	Update(cxt context.Context, user *model.UserM) error
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

func (u *users) Get(cxt context.Context, userInput *model.UserM) (*model.UserM, error) {
	var userOutput model.UserM
	if err := u.db.Where(&userInput).First(&userOutput).Error; err != nil {
		return nil, err
	}
	return &userOutput, nil
}

func (u *users) Update(cxt context.Context, user *model.UserM) error {
	return u.db.Save(&user).Error
}
