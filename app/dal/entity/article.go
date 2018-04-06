package entity

type Article struct {
	BaseEntity
	ArticleTitle		string
	ArticleFrom			string
	ArticleUrl			string
	InternelUrl			string
	ArticleLabels 		string
	ArticleKeywords		string
	FavoriteTimes		int
	ViewTimes			int
	MarkTimes			int
	IsTechnology		string
	IsBlockchain		string
	State 				string
	Comment 			string
}
