package entity

type ArticleSource struct {
	BaseEntity  `xorm:"extends"`
	SourceName string `json:"source_name"`
	SourceType string `json:"source_type"`
	SourceUrl  string `json:"source_url"`
}
