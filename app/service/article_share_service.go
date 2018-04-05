package service

import (
	"github.com/geekappio/itonchain/app/common/model/dal"
)

type ArticleShareService interface {
	DoArticleShare(share *dal.ArticleShare) bool
}

func NewArticleShareService() ArticleShareService {
	return ArticleShareServiceImpl{}
}

type ArticleShareServiceImpl struct {
}

func (self ArticleShareServiceImpl) DoArticleShare(model *dal.ArticleShare) bool {
	println("模拟文件分享完成")
	return true
}