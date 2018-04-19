package service

import (
	"strconv"
	"strings"

	"github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/dal/dao"
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"github.com/geekappio/itonchain/app/util"

	"time"

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
	wechatUser := entity.WechatUser{
		BaseEntity: entity.BaseEntity{
			GmtCreate: time.Now(),
			GmtUpdate: time.Now(),
		},
		IsDel: "NO",
	}
	copier.Copy(wechatUser, request)
	//查询openId是否存在，存在报错
	wechatUserSqlMapper := dao.GetWechatUserSqlMapper(nil)
	user, err := wechatUserSqlMapper.SelectUser(request.OpenId)
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
		return model.NewFailedResponseModel(enum.DB_INSERT_ERROR, "添加用户数据失败")
	} else {
		return model.NewFailedResponseModelWithData(enum.SYSTEM_SUCCESS, "用户注册成功",
			id)
	}
}

func (self *WechatUserService) FindUserByOpenId(openId string) (wechatUser *entity.WechatUser, err error) {
	wechatUserSqlMapper := dao.GetWechatUserSqlMapper(nil)
	user, err := wechatUserSqlMapper.SelectUser(openId)
	if err != nil {
		util.LogError("根据openId查询用户失败 ", openId, user, err)
	}
	return user, err
}

func (service *WechatUserService) ChangingArticleCategoryOrder(request *model.ArticleCategoryOrderChangeRequest) *model.ResponseModel {
	// Here calls dao method to access database.
	userModel, err := dao.GetWechatUserSqlMapper(nil).SelectUser(request.OpenId)
	if err != nil {
		util.LogError("Error happened when getting user model from wechat_user table with openId: ", request.OpenId, err)
		return model.NewFailedResponseModel(enum.DB_ERROR, "从数据库查询数据发送错误")
	}
	if userModel == nil {
		util.LogInfo("Cannot find user by specified open id:", request.OpenId)
		return model.NewFailedResponseModel(enum.USER_NOT_EXISTS, "指定用户不存在")
	}

	orders := userModel.CategoryOrders
	if orders == "" {
		return model.NewFailedResponseModel(enum.NULL_CATEGORY_ORDERS, "空的目录顺序项")
	} else {
		categoryStr := strconv.FormatInt(request.CategoryId, 10)
		categories := strings.Split(userModel.CategoryOrders, config.FIELD_CATEGORY_ORDRES_SEPARATER)
		for index, v:= range categories {
			if v == categoryStr {
				enumValue := field_enum.ValueOf(v)
				if enumValue == field_enum.UP {
					switch {
					case index == 0:
						// 已经是第一个，无法上移
						return model.NewFailedResponseModel(enum.IS_FIRST_CATEGORY, "已经是第一个目录项，无法上移")

					case index > 0 && index < len(categories):
						// 可以移动
						categories = util.StrigArrayRemove(categories, v)
						categories = util.StringArrayInsert(categories, index - 1, v)

						_, updateErr := dao.GetWechatUserSqlMapper(nil).UpdateCategoryOrders(request.OpenId, strings.Join(categories, config.FIELD_CATEGORY_ORDRES_SEPARATER))
						if updateErr != nil {
							return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR, "更新目录项数据失败")
						}

					case index == -1 || index > len(categories):
						// 没找到
						return model.NewFailedResponseModel(enum.NOT_FIND_SPECIFIED_CATEGORY, "没有发现指定的目录项，category：" + strconv.FormatInt(request.CategoryId, 10))
					}

				} else if enumValue == field_enum.DOWN {

					switch {
					case index == len(categories) - 1:
						// 已经是最后一个，无法下移
						return model.NewFailedResponseModel(enum.IS_LAST_CATEGORY, "已经是最后一个目录项，无法下移")

					case index >= 0 && index < len(categories) - 1:
						// 可以移动
						categories = util.StringArrayInsert(categories, index + 2, v)
						categories = util.StringArrayRemoveByIndex(categories, index)

						_, updateErr := dao.GetWechatUserSqlMapper(nil).UpdateCategoryOrders(request.OpenId, strings.Join(categories, config.FIELD_CATEGORY_ORDRES_SEPARATER))
						if updateErr != nil {
							return model.NewFailedResponseModel(enum.DB_UPDATE_ERROR, "更新目录项数据失败")
						}

					case index == -1 || index > len(categories):
						// 没找到
						return model.NewFailedResponseModel(enum.NOT_FIND_SPECIFIED_CATEGORY, "没有发现指定的目录项，category：" + strconv.FormatInt(request.CategoryId, 10))
					}
				} else {
					return model.NewFailedResponseModel(enum.INVALID_REQUEST_FIELD_VALUE, "错误的参数值，UpDown：" + request.UpDown)
				}
			}
		}

		return model.NewSuccessResponseModel()
	}
}
