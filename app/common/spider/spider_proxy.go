package spider

import (
	"github.com/geekappio/itonchain/app/dal/entity"
	"github.com/geekappio/itonchain/app/model/field_enum"
	"errors"
	"github.com/geekappio/itonchain/app/common/spider/impl"
)

type Spider interface {
	Start()
	Exit()
	AddSource(source string) bool
}

var (
	feedSpider = impl.NewFeedSpider()
)


// func init() {
// 	feedSpider.Start()
// }

func Capture(sources []*entity.ArticleSource) error {
	for _, source := range sources {
		switch field_enum.ValueOf(source.SourceType) {
		case field_enum.FEED:
			feedSpider.AddSource(source.SourceUrl)
		default:
			return errors.New("不支持的文章源种类")
		}
	}
	return nil
}