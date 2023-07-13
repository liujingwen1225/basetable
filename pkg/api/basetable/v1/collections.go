package v1

type CreateCollectionsRequest struct {
	Name    string                            `json:"name" valid:"required,alphanum,max=50,min=1"`
	Type    int                               `json:"type" valid:"required"`
	Options map[string]any                    `json:"options"`
	Fields  []*CreateCollectionsFieldsRequest `json:"fields" valid:"required,gte=1,dive"`
}

type CreateCollectionsFieldsRequest struct {
	Name     string `json:"name" valid:"required"`
	Type     int    `json:"type" valid:"required"`
	Required int    `json:"required" valid:"required"`
	Options  string `json:"options"`
}

type CollectionsResponse struct {
	ID          int                              `json:"id"`
	Name        string                           `json:"name"`
	Type        int                              `json:"type"`
	SourceTable string                           `json:"source_table"`
	Options     map[string]any                   `json:"options"`
	Fields      []CreateCollectionsFieldsRequest `json:"fields"`
}

type CollectionsFieldsResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Required int    `json:"required"`
	System   int    `json:"system"`
	Options  string `json:"options"`
}
