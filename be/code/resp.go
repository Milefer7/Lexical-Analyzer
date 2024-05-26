package code

import "github.com/gin-gonic/gin"

// 定义一个响应状态码的类型
type ResCode int

const EmptyData = ""

// 业务逻辑状态码
const (
	// Success code为0表示成功
	Success ResCode = 1
	Fail    ResCode = 0
)

// 业务逻辑状态信息描述
var recodeText = map[ResCode]string{
	Success: "处理成功",
	Fail:    "处理失败",
}

// RespMsg : 响应数据结构
type RespMsg struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// StatusText 返回状态码的文本。如果代码为空或未知状态码则返回error
func (code ResCode) StatusText() string {
	msg, ok := recodeText[code]
	if ok {
		return msg
	} else {
		return "未知状态码"
	}
}

// NewRespMsg : 生成response对象
func NewRespMsg(code ResCode, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  code.StatusText() + " " + msg,
		Data: data,
	}
}

// CommonResp : 通用响应
// 传入参数：1. gin.Context 2. http状态码 3. 自定义状态码 4. 错误信息（如果正确就传空） 5. 数据（要返回是数据）
func CommonResp(c *gin.Context, httpCode int, statusCode ResCode, msg string, data interface{}) {
	c.JSON(httpCode, *NewRespMsg(statusCode, msg, data))
	c.Abort() //此路由后的 gin.HandlerFunc 将不再被调用
}
