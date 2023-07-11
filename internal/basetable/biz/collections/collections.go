package collections

import (
	"basetable.com/internal/basetable/store"
	"basetable.com/internal/pkg/model"
	v1 "basetable.com/pkg/api/basetable/v1"
	"context"
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
	var collections *model.CollectionsM
	var collectionsFields *model.CollectionsFieldsM
	err := query.CollectionsM.Create(collections)

	return err
}
