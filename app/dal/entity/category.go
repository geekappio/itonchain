package entity

type Category struct {
	BaseEntity
	UserId 			int
	CategoryName 	string
	Description 	string
	ArticleCount 	int
	IsDel 			string
}
