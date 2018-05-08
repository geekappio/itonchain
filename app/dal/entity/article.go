package entity

type Article struct {
	BaseEntity             `xorm:"extends"`
	ArticleTitle    string `xorm:'article_title' varchar(1000) notnull`
	ArticleFrom     string `xorm:'article_from' varchar(255)`
	ArticleUrl      string `xorm:'article_url' varchar(1000) notnull`
	InternelFid     string `xorm:'internel_fid' varchar(100) notnull`
	InternelUrl     string `xorm:'internel_url' varchar(1000) notnull`
	InternelSize    int64  `xorm:'internel_size' bigint(10) default(0)`
	ContentType     string `xorm:'content_type' varchar(20) notnull`
	Images          string `xorm:'images' varchar(4096)`
	PreviewLayout   string `xorm:'preview_layout' varchar(20)`
	ArticleLabels   string `xorm:'article_labels' varchar(1000)`
	ArticleKeywords string `xorm:'article_keywords' varchar(1000)`
	FavoriteTimes   int32  `xorm:'favorite_times' bigint(10) default(0)`
	ViewTimes       int32  `xorm:'view_times' bigint(10) default(0)`
	MarkTimes       int32  `xorm:'mark_times' bigint(10) default(0)`
	IsTechnology    string `xorm:'is_technology' notnull default('NO')`
	IsBlockchain    string `xorm:'is_blockchain' notnull unqiue default('NO')`
	State           string `xorm:'state' varchar(45)`
	Comment         string `xorm:'comment' varchar(200)`
}
