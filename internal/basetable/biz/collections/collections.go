package collections

import (
	"context"
	"strings"

	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	"basetable.com/pkg/api"
	v1 "basetable.com/pkg/api/basetable/v1"
	"basetable.com/pkg/util"
	"github.com/jinzhu/copier"
)

type ICollectionsBiz interface {
	Create(cxt context.Context, request *v1.CreateCollectionsRequest) error
	List(cxt context.Context, page *api.PageRequest) ([]*model.CollectionsM, error)
	Update(cxt context.Context, request *v1.CreateCollectionsRequest) error
	Delete(cxt context.Context, id int) error
}

var (
	_     ICollectionsBiz = &Collections{}
	query                 = store.Q
)

type Collections struct {
}

func (c *Collections) Delete(cxt context.Context, id int) error {
	var collections = model.CollectionsM{BaseModelM: model.BaseModelM{ID: id}}
	if _, err := query.CollectionsM.Delete(&collections); err != nil {
		return err
	}
	return nil
}

func (c *Collections) Update(cxt context.Context, request *v1.CreateCollectionsRequest) error {
	var collections model.CollectionsM
	_ = copier.Copy(&collections, &request)
	// type and fieldName check
	var fieldNames = make(map[string]int)
	if !collections.IsInCollectionTypes() {
		return errno.ErrCollectionsTypeNotFound
	}
	for _, fields := range collections.Fields {
		if !fields.IsInFieldType() {
			return errno.ErrCollectionsFieldsTypeNotFound
		}
		_, ok := fieldNames[fields.Name]
		if !ok {
			fieldNames[fields.Name] = 1
		} else {
			return errno.ErrCollectionsFieldNameDuplicate
		}
	}
	// Collections name check (case insensitive)
	if result, _ := query.CollectionsM.CollectionNameByLowercase(strings.ToLower(collections.Name)); len(result) > 0 {
		if len(result) > 2 || result[0].ID != request.ID {
			return errno.ErrCollectionsExist
		}
	}
	// update
	if err := query.Transaction(func(tx *store.Query) error {
		if _, err := tx.CollectionsM.Updates(&collections); err != nil {
			return errno.ErrCollectionsUpdate
		}
		// todo change table options

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Collections) Create(cxt context.Context, request *v1.CreateCollectionsRequest) error {
	var collections model.CollectionsM
	_ = copier.Copy(&collections, &request)
	// type and fieldName check
	var fieldNames = make(map[string]int)
	if !collections.IsInCollectionTypes() {
		return errno.ErrCollectionsTypeNotFound
	}
	for _, fields := range collections.Fields {
		if !fields.IsInFieldType() {
			return errno.ErrCollectionsFieldsTypeNotFound
		}
		_, ok := fieldNames[fields.Name]
		if !ok {
			fieldNames[fields.Name] = 1
		} else {
			return errno.ErrCollectionsFieldNameDuplicate
		}
	}
	// Collections name check (case insensitive)
	if result, _ := query.CollectionsM.CollectionNameByLowercase(strings.ToLower(collections.Name)); len(result) > 0 {
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

func NewCollectionsBiz() *Collections {
	return &Collections{}
}
