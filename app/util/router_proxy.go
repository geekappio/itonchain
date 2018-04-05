package util

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"encoding/json"
	"net/http"
	"github.com/geekappio/itonchain/app/common/model/api"
	"github.com/geekappio/itonchain/app/common/model/enum"
)

// FIXME 目前该方法为实验性质，传入 handler 必须遵循规则为 func(BaseRequest) (BaseResponse, errorCode string)
func AddPostRouter(engine *gin.Engine, path string, handler interface{}) {
	typ := reflect.TypeOf(handler)
	val := reflect.ValueOf(handler)
	if typ.Kind() != reflect.Func || typ.NumIn() != 1 || typ.NumOut() != 2 {
		panic("传入 handler 必须遵循规则为 func(BaseRequest) (BaseResponse, errorCode string)")
	}
	engine.POST(path, func(c *gin.Context) {
		rawReq, err := c.GetRawData()
		if err != nil {
			LogError(err)
			return
		}
		req := reflect.New(typ.In(0)).Elem()
		json.Unmarshal(rawReq, req.Addr().Interface())
		results := val.Call([]reflect.Value{req})

		response := results[0].Interface()
		errorCode := results[1].Interface().(string)
		head := api.ResponseHead{
			ReturnCode:errorCode,
			ReturnMsg:enum.GetErrorMsg(errorCode),
			ReturnData:response,
		}
		c.JSON(http.StatusOK, head)
	})
}
