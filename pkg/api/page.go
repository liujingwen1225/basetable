package api

type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
