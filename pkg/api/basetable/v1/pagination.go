package v1

type Pagination struct {
	Page     int `json:"page" valid:"number"`
	PageSize int `json:"PageSize" valid:"number"`
}

func (p *Pagination) GetPage() (offset, limit int) {
	return (p.Page - 1) * p.PageSize, p.PageSize
}
