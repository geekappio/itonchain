package web

import (
	"github.com/gin-gonic/gin"
	"github.com/geekappio/itonchain/app/service"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/util"
	"net/http"
	"strconv"
	"github.com/geekappio/itonchain/app/common/seaweedfs"
)

// pending 分页列表函数
func ArticlePendingList(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	mapper := service.GetArticlePendingService()

	count, _ := mapper.GetArticlePendingCount()
	totalPages := (count + pageSize - 1) / pageSize
	if pageSize > totalPages || pageSize < 1 {
		pageSize = 1	// 校验并设置默认页数，应该给提示，但是返回显示比较麻烦
	}
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

// FIXME pending 保存函数
func ArticlePendingSave(c *gin.Context) {
	// 获取原始Pending的fid 和 更新的内容
	fid := c.PostForm("fid")
	content := c.PostForm("content")

	// 调用fs函数更新文件
	_, err := seaweedfs.UpdateResourceContent("", fid, []byte(content), nil, true)

	// 返回成功/失败
	c.String(http.StatusOK, err.Error())
}
