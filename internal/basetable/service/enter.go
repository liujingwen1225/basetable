package service

type ServiceGroup struct {
	UserService
	collectionsBiz
}

var ServiceGroupApp = new(ServiceGroup)
