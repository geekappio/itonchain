package entity

type ArticleFavorite struct {
	BaseEntity
	ArticleId 			int64 `xorm:'article_id' bigint(20) notnull`
	UserId 				int64 `xorm:'user_id' bigint(20) notnull`
}
