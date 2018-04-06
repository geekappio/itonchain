package service

import (
	"github.com/geekappio/itonchain/app/dal/entity"
)

type ArticleShareService interface {
	DoArticleShare(share *entity.ArticleShare) bool
}

func NewArticleShareService() ArticleShareService {
	return ArticleShareServiceImpl{}
}

type ArticleShareServiceImpl struct {
}

func (self ArticleShareServiceImpl) DoArticleShare(model *entity.ArticleShare) bool {
	println("模拟文件分享完成")
	return true
}