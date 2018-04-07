package service

import (
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/util"
	"github.com/geekappio/itonchain/app/model"

	"github.com/geekappio/itonchain/app/enum"
	"github.com/jinzhu/copier"
)

var wechatUserService *WechatUserService

func GetWechatUserService() *WechatUserService {
	if wechatUserService == nil {
		wechatUserService = &WechatUserService{}
	}

	return wechatUserService
}

type WechatUserService struct {
}

func (self *WechatUserService) CreateUser(request *model.WechatUserRequest) (*model.ResponseModel) {
	wechatUser := entity.WechatUser{}
	copier.Copy(wechatUser, request)
	//查询openId是否存在，存在报错
	wechatUserSqlMapper := dao.GetWechatUserSqlMapper()
	bool, err := wechatUserSqlMapper.UserRegister(&wechatUser)
	if err != nil {
		util.LogError(err)
	}
	if bool {
		//openId已存在
		util.LogError("用户已存在")
	}
	//否则创建用户
	id, err := wechatUserSqlMapper.InsertUser(&wechatUser)

	if err != nil {
		util.LogError("Error happened when inserting wechat_user: ", wechatUser, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR,
			ReturnMsg:  "添加category数据失败",
		}
	} else {
		return &model.ResponseModel{
			ReturnCode: enum.SYSTEM_SUCCESS,
			ReturnMsg:  "用户注册成功",
			ReturnData: id,
		}
	}
}

func (self *WechatUserService) FindUserByOpenId(openId string) *entity.WechatUser {
	return &entity.WechatUser{
		BaseEntity: entity.BaseEntity{Id: 123321,},
	}
}

func (service *WechatUserService) ChangingArticleCategoryOrder(request *model.ArticleCategoryOrderChangeRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	// TODO ...
	return &model.ResponseModel{}
}
