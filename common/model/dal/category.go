package dal

type Category struct {
	Base
	UserId 			int
	CategoryName 	string
	Description 	string
	ArticleCount 	int
	IsDel 			string
}
