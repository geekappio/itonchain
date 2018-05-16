package web

import (
	"github.com/gin-gonic/gin"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"net/http"
)

func ArticlePendingList(c *gin.Context) {
	articlePendings, _ := service.GetArticlePendingService().GetArticlePendingList(1, 50, "")
	articles := make([]*model.ArticlePendingModel, len(articlePendings))
	for i, articlePending := range articlePendings {
		articles[i] = &model.ArticlePendingModel{
			ArticlePendinId: articlePending.Id,
			ArticleTitle:    articlePending.ArticleTitle,
			ArticleFrom:     articlePending.ArticleFrom,
			ArticleUrl:      articlePending.ArticleUrl,
			InternelFid:     articlePending.InternelFid,
			InternelUrl:     articlePending.InternelUrl,
			InternelSize:    articlePending.InternelSize,
			ArticleKeywords: articlePending.ArticleKeywords,
			GmtCreate:       util.TimeFormat(articlePending.GmtCreate),
			GmtUpdate:       util.TimeFormat(articlePending.GmtUpdate),
		}
	}
	c.HTML(http.StatusOK, "article-list.html", articles)
}