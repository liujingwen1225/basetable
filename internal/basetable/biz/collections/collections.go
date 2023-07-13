package collections

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/errno"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"context"
	"github.com/jinzhu/copier"
)

type ICollectionsBiz interface {
	Create(cxt context.Context, request *v1.CreateCollectionsRequest) error
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
	var collectionsFields = make([]*model.CollectionsFieldsM, len(request.Fields))

	_ = copier.Copy(&collections, &request)
	_ = copier.Copy(&collectionsFields, &request.Fields)

	err := query.Transaction(func(tx *store.Query) error {
		if err := tx.CollectionsM.Create(&collections); err == nil {
			for _, field := range collectionsFields {
				field.CollectionsID = collections.ID
			}
			if ferr := tx.CollectionsFieldsM.CreateInBatches(collectionsFields, 1000); ferr != nil {
				return errno.ErrCollectionsCreate
			}
		} else {
			return errno.ErrCollectionsCreate
		}
		//todo create table

		return nil
	})
	return err
}
