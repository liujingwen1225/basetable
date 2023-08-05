package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *datastore
)

type IStore interface {
	Users() UserStore
	Collections() CollectionsStore
	Schema() SchemaStore
}

var _ IStore = &datastore{}

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Collections() CollectionsStore {
	return newCollectionsStore(ds.db)
}

func (ds *datastore) Schema() SchemaStore {
	return newSchemaStore(ds.db)
}

func (ds *datastore) Users() UserStore {
	return newUserStore(ds.db)
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db: db}
	})
	return S
}
