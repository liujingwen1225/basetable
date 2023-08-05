package service

// import (
//
//	"basetable.com/internal/pkg/model"
//	v1 "basetable.com/pkg/api/basetable/v1"
//	"github.com/jinzhu/copier"
//
// )
type collectionsBiz struct {
}

//
//func (c *collectionsBiz) Create(request *v1.CollectionsRequest) (*v1.CollectionsResponse, error) {
//	var collectionM *model.CollectionsM
//	var schemaMs []*model.SchemaM
//	_ = copier.Copy(&collectionM, request)
//	_ = copier.Copy(&schemaMs, request.Schemas)
//}
//
//func (c *collectionsBiz) Deleted(ids []int) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *collectionsBiz) GetById(id int) (*v1.CollectionsResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *collectionsBiz) List(request *v1.ListUserRequest) ([]*v1.CollectionsResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (c *collectionsBiz) Update(request *v1.UpdateUserRequest) (*v1.CollectionsResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}
