// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package store

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                  = new(Query)
	CollectionsFieldsM *collectionsFieldsM
	CollectionsM       *collectionsM
	UserM              *userM
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CollectionsFieldsM = &Q.CollectionsFieldsM
	CollectionsM = &Q.CollectionsM
	UserM = &Q.UserM
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		CollectionsFieldsM: newCollectionsFieldsM(db, opts...),
		CollectionsM:       newCollectionsM(db, opts...),
		UserM:              newUserM(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CollectionsFieldsM collectionsFieldsM
	CollectionsM       collectionsM
	UserM              userM
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CollectionsFieldsM: q.CollectionsFieldsM.clone(db),
		CollectionsM:       q.CollectionsM.clone(db),
		UserM:              q.UserM.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CollectionsFieldsM: q.CollectionsFieldsM.replaceDB(db),
		CollectionsM:       q.CollectionsM.replaceDB(db),
		UserM:              q.UserM.replaceDB(db),
	}
}

type queryCtx struct {
	CollectionsFieldsM ICollectionsFieldsMDo
	CollectionsM       ICollectionsMDo
	UserM              IUserMDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CollectionsFieldsM: q.CollectionsFieldsM.WithContext(ctx),
		CollectionsM:       q.CollectionsM.WithContext(ctx),
		UserM:              q.UserM.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
