package entity

type Category struct {
	BaseEntity
	UserId       int64  `xorm:'user_id' bigint(20) notnull`
	CategoryName string `xorm:'category_name' notnull varchar(45)`
	Description  string `xorm:'description' varchar(500)`
	ArticleCount int32  `xorm:'article_count' bigint(10)`
	IsDel        string `xorm:'is_del' notnull default('NO')`
}
