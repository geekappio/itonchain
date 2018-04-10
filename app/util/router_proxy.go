package util

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"github.com/geekappio/itonchain/app/model"
	"github.com/geekappio/itonchain/app/enum"
)

// FIXME 目前该方法为实验性质，传入 handler 必须遵循规则为 func(BaseRequest) (BaseResponse, ErrorCode)
// FIXME handler 目前有两种方案，一种直接返回ResponseMode, 另一种返回BaseResponse,ErrorCode，暂且做下兼容使编译通过
func AddPostRouter(engine *gin.Engine, path string, handler interface{}) {
	typ := reflect.TypeOf(handler)
	val := reflect.ValueOf(handler)
	if typ.Kind() != reflect.Func || typ.NumIn() != 1 {
		panic("传入 handler 必须遵循规则为 func(BaseRequest) (BaseResponse, ErrorCode)")
	}
	engine.POST(path, func(c *gin.Context) {
		rawReq, err := c.GetRawData()
		if err != nil {
			LogError("Error happened when getting data from http request body.", err)
			return
		}

		// 记录Request日志
		LogRequest(path, rawReq)

		// 解码request body到request model对象
		req := reflect.New(typ.In(0)).Elem()
		err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(rawReq, req.Addr().Interface())
		if err != nil {
			LogError("Error happened when decoding request data.", err)
		}

		// 利用反射机制执行Service Call
		results := val.Call([]reflect.Value{req})
		response := results[0].Interface()

		// Record response log
		LogResponse(path, &response)

		if 1 == len(results) {
			c.JSON(http.StatusOK, response)
		} else {
			errorCode := results[1].Interface().(enum.ErrorCode)
			c.JSON(http.StatusOK, &model.ResponseModel{
				ReturnCode: errorCode.GetRespCode(),
				ReturnMsg: errorCode.GetRespMsg(),
				ReturnData: response,
			})
		}
	})
}
