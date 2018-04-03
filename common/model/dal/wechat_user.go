package dal

type WechatUser struct {
	Base
	OpenId 			string
	NickName 		string
	AvatarUrl 		string
	Gender 			string
	City			string
	Province 		string
	Country 		string
	Language 		string
	IdDel 			string
	CategoryOrders 	string
}
