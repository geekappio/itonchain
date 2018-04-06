package service

import (
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model"
)

var wechatUserService *WechatUserService

// GetArticleCategoryService returns ArticleCategory service instance which provides method calls.
func GetWechatUserService() *WechatUserService {
	if wechatUserService == nil {
		wechatUserService = &WechatUserService{}
	}

	return wechatUserService
}


// Implementation struct of ArticleCategory to bind functions wi
type WechatUserService struct {
}

func (self *WechatUserService) CreateUser(model *entity.WechatUser) bool {
	println("模拟创建用户成功！")
	return true
}

func (self *WechatUserService) FindUserByOpenId(openId string) *entity.WechatUser {
	return &entity.WechatUser{
		BaseEntity: entity.BaseEntity{Id: 123321,},
	}
}

func (service *WechatUserService) ChaningArticleCategoryOrder(request *model.ArticleCategoryOrderChangeRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	// TODO ...
	return &model.ResponseModel{}
}