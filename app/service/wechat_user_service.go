package service

import (
	"github.com/geekappio/itonchain/app/common/model/dal"
)

type WechatUserService interface {
	CreateUser(model *dal.WechatUser) bool
	FindUserByOpenId(openId string) *dal.WechatUser
}

func NewWechatUserService() WechatUserService {
	return WechatUserServiceImpl{}
}

type WechatUserServiceImpl struct {
}

func (self WechatUserServiceImpl) CreateUser(model *dal.WechatUser) bool {
	println("模拟创建用户成功！")
	return true
}

func (self WechatUserServiceImpl) FindUserByOpenId(openId string) *dal.WechatUser {
	return &dal.WechatUser{
		Base:dal.Base{Id:123321,},
	}
}