package api

type ArticleShareRequest struct {
	BaseRequest
	OpenId 		string		`json:"openId"`
	ArticleId 	int			`json:"articleId"`
}

type ArticleShareResponse struct {
	BaseResponse
	ShareTimes 		int		`json:"shareTimes"`
}
