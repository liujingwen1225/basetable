package store

import (
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"context"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(cxt context.Context, user *model.UserM) error
	List(cxt context.Context, user *model.UserM, pagination v1.Pagination) ([]*model.UserM, error)
	GetById(cxt context.Context, id int) (*model.UserM, error)
	GetByUserName(cxt context.Context, name string) (*model.UserM, error)
	Update(cxt context.Context, user *model.UserM) (*model.UserM, error)
	Delete(cxt context.Context, ids []int) error
}

var _ UserStore = &users{}

type users struct {
	db *gorm.DB
}

func (u *users) GetByUserName(cxt context.Context, name string) (*model.UserM, error) {
	var res model.UserM
	if err := u.db.Where("username = ?", name).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *users) List(cxt context.Context, user *model.UserM, pagination v1.Pagination) ([]*model.UserM, error) {
	var res []*model.UserM
	query := u.db
	if user.Username != "" {
		query = query.Where("username LIKE ?", "%"+user.Username+"%")
	}
	if user.Nickname != "" {
		query = query.Where("nickname LIKE ?", "%"+user.Nickname+"%")
	}
	offset, limit := pagination.GetPage()
	tx := query.Offset(offset).Limit(limit).Find(&res)

	if tx.Error != nil {
		return nil, errno.ErrUserNotFound
	}
	return res, nil
}

func (u *users) GetById(cxt context.Context, id int) (*model.UserM, error) {
	var userOutput model.UserM
	if err := u.db.Where("id = ?", id).First(&userOutput).Error; err != nil {
		return nil, err
	}
	return &userOutput, nil
}

func newUserStore(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(cxt context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func (u *users) Update(cxt context.Context, user *model.UserM) (*model.UserM, error) {
	return user, u.db.Omit("password", "created_at").Save(&user).Error
}
func (u *users) Delete(cxt context.Context, ids []int) error {
	return u.db.Delete(&model.UserM{}, ids).Error
}
