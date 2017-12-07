package base

// 自定义错误返回内容
import (
	"github.com/astaxie/beego"
)

var errorBody = make(map[int][]byte, 0)

func init() {
	errorBody[401] = []byte("unAuthorized \n")
	errorBody[404] = []byte("request not found\n")
	errorBody[501] = []byte("server error\n")
}

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.Ctx.ResponseWriter.WriteHeader(401)
	c.Ctx.ResponseWriter.Write(errorBody[401])
}

func (c *ErrorController) Error404() {
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.Ctx.ResponseWriter.Write(errorBody[404])
}

func (c *ErrorController) Error501() {
	c.Ctx.ResponseWriter.WriteHeader(501)
	c.Ctx.ResponseWriter.Write(errorBody[501])
}

// 用于获取页面发生错误信息时的返回结构
type ErrorData struct {
	ErrMsg string
	Data   interface{}
}

func NewErrorData(msg string) *ErrorData {
	return &ErrorData{
		ErrMsg: msg,
	}
}
