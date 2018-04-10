package entity

type ArticlePending struct {
	BaseEntity             `xorm:"extends"`
	ArticleTitle    string `xorm:'article_title' varchar(255) notnull`
	ArticleFrom     string `xorm:'article_from' varchar(255) notnull`
	ArticleUrl      string `xorm:'article_url' varchar(255) notnull `
	ArticleKeywords string `xorm:'article_keywords' varchar(1000)`
}
