package entity

type WechatUser struct {
	BaseEntity
	OpenId 			string
	NickName 		string
	AvatarUrl 		string
	Gender 			string
	City			string
	Province 		string
	Country 		string
	Language 		string
	IsDel 			string
	CategoryOrders 	string
}
