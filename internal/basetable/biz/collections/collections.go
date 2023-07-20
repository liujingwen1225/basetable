package collections

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	"basetable.com/pkg/api"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/util"
	"context"
	"github.com/jinzhu/copier"
	"strings"
)

type ICollectionsBiz interface {
	Create(cxt context.Context, request *v1.CreateCollectionsRequest) error
	List(cxt context.Context, page *api.PageRequest) ([]*model.CollectionsM, error)
}

var (
	_     ICollectionsBiz = &Collections{}
	query                 = store.Q
)

type Collections struct {
}

func NewCollectionsBiz() *Collections {
	return &Collections{}
}

func (c *Collections) Create(cxt context.Context, request *v1.CreateCollectionsRequest) error {
	var collections model.CollectionsM
	_ = copier.Copy(&collections, &request)
	// type check
	if !collections.IsInCollectionTypes() {
		return errno.ErrCollectionsTypeNotFound
	}
	for _, fields := range collections.Fields {
		if !fields.IsInFieldType() {
			return errno.ErrCollectionsTypeNotFound
		}
	}
	// name check (case insensitive)
	if result, _ := query.CollectionsM.CollectionNameByLowercase(strings.ToLower(collections.Name)); result > 0 {
		return errno.ErrCollectionsExist
	}
	// system fields init
	collections.CollectionSystemFieldsInit()

	// options init

	// create
	if err := query.Transaction(func(tx *store.Query) error {
		if err := tx.CollectionsM.Create(&collections); err != nil {
			return errno.ErrCollectionsCreate
		}
		// todo create table options

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Collections) List(cxt context.Context, page *api.PageRequest) ([]*model.CollectionsM, error) {
	offset, limit := util.Pagination(page)
	result, _, err := query.CollectionsM.Preload(query.CollectionsM.Fields).FindByPage(offset, limit)
	return result, err
}
