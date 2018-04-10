package service

import (
	"strings"

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
	wechatUserSqlMapper := dao.GetWechatUserSqlMapper(nil)
	user ,err := wechatUserSqlMapper.SelectUser(request.OpenId)
	if err != nil {
		util.LogError(err)
	}
	if user != nil {
		//openId已存在
		util.LogError("用户已存在")
	}
	//否则创建用户
	id, err := wechatUserSqlMapper.InsertUser(&wechatUser)

	if err != nil {
		util.LogError("Error happened when inserting wechat_user: ", wechatUser, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_INSERT_ERROR.GetRespCode(),
			ReturnMsg: "添加category数据失败",
		}
	} else {
		return &model.ResponseModel{
			ReturnCode: enum.SYSTEM_SUCCESS.GetRespCode(),
			ReturnMsg: "用户注册成功",
			ReturnData: id,
		}
	}
}

func (self *WechatUserService) FindUserByOpenId(openId string) (wechatUser *entity.WechatUser, err error) {
	wechatUserSqlMapper := dao.GetWechatUserSqlMapper(nil)
	user, err := wechatUserSqlMapper.SelectUser(openId)
	if err != nil {
		util.LogError("根据openId查询用户失败 ", user, err)
	}
	return user, err
}

func (service *WechatUserService) ChangingArticleCategoryOrder(request *model.ArticleCategoryOrderChangeRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	userModel, err := dao.GetWechatUserSqlMapper().SelectUser(request.OpenId)
	if err != nil{
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, err)
		return &model.ResponseModel{
			ReturnCode: enum.DB_ERROR.GetRespCode(),
			ReturnMsg:  "从数据库查询数据发送错误",
		}
	}
	if userModel == nil {
		util.LogInfo("Cannot find user by specified open id:", request.OpenId)
		return &model.ResponseModel{
			ReturnCode: enum.USER_NOT_EXISTS.GetRespCode(),
			ReturnMsg:  "指定用户不存在",
		}
	}

	orders := userModel.CategoryOrders
	if orders == "" {
		return &model.ResponseModel{
			ReturnCode: enum.NULL_CATEGORY_ORDERS.GetRespCode(),
			ReturnMsg:  "空的目录顺序项",
		}
	} else {
		// TODO, HENRY, 20180409, 根据参数调整次序
		strings.Split(orders, ",")
	}

	dao.GetWechatUserSqlMapper(nil).UpdateCategoryOrders(request.OpenId, orders)
	return &model.ResponseModel{}
}
