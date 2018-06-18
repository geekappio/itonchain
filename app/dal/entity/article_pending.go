package entity

type ArticlePending struct {
	BaseEntity             `xorm:"extends"`
	ArticleTitle    string `xorm:'article_title' varchar(255) notnull`
	ArticleFrom     string `xorm:'article_from' varchar(255) notnull`
	ArticleUrl      string `xorm:'article_url' varchar(1000) notnull`
	InternelFid     string `xorm:'internel_fid' varchar(100) notnull`
	InternelUrl     string `xorm:'internel_url' varchar(1000) notnull`
	InternelSize    int64  `xorm:'internel_size' bigint(10) default(0)`
	ArticleKeywords string `xorm:'article_keywords' varchar(1000)`
	State           string `xorm:'state' varchar(20)`
}
