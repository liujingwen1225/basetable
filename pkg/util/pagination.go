package util

import "basetable.com/pkg/api"

func Pagination(request *api.PageRequest) (offset, limit int) {
	return (request.Page - 1) * request.PageSize, request.PageSize
}
