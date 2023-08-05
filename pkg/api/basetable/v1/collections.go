package v1

type CollectionsRequest struct {
	Name    string          `valid:"required;alphanum;gt=30" json:"name"`
	Type    int             `valid:"required" json:"type"`
	Options *string         `valid:"" json:"options"`
	Schemas []SchemaRequest `valid:"required;desc" json:"schemas"`
}

type CollectionsResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	SourceTable string  `json:"source_table"`
	Options     *string `json:"options"`
}
