package web

import (
	"github.com/gin-gonic/gin"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"net/http"
	"strconv"
)

func ArticlePendingList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "50"))

	mapper := service.GetArticlePendingService()

	count, _ := mapper.GetArticlePendingCount()
	totalPages := (count + pageSize - 1) / pageSize
	articlePendings, _ := mapper.GetArticlePendingList(pageNum, pageSize, "")
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
	c.HTML(http.StatusOK, "article-list.html", gin.H{
		"Articles": articles,
		"CurPage": pageNum,
		"PageSize": pageSize,
		"TotalPages": totalPages,
	})
}