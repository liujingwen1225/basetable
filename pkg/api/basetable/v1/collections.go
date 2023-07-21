package v1

type CreateCollectionsRequest struct {
	ID      int                               `json:"id"`
	Name    string                            `json:"name" valid:"required,alphanum,max=50,min=1"`
	Type    string                            `json:"type" valid:"required"`
	Index   string                            `json:"index"`
	Options string                            `json:"options"`
	Fields  []*CreateCollectionsFieldsRequest `json:"fields" valid:"required,gte=1,dive"`
}

type CreateCollectionsFieldsRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" valid:"required"`
	Type     string `json:"type" valid:"required"`
	Required *bool  `json:"required" valid:"required"`
	Options  string `json:"options"`
}
