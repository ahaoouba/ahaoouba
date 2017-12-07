package smsutil

import "fmt"
import "strings"
import "encoding/json"

type UserInfo struct {
	// 用户账号：长度最大6个字符，统一大写,如提交参数中包含apikey，
	// 则可以不用填写该参数及pwd，两种鉴权方式中只能选择一种方式来
	Userid string `json:"userid"`

	// 用户密码：定长小写32位字符, 如提交参数中包含apikey，则可以
	// 不用填写该参数及userid，两种鉴权方式中只能选择一种方式来进
	// 行鉴权。
	Pwd string `json:"pwd"`

	// 时间戳,格式为:MMDDHHMMSS,即月日时分秒,定长10位,月日时分秒不足2位时左补0.时间戳请获取您真实的服务器时间,不要填写固定的时间,否则pwd参数起不到加密作用
	Timestamp string `json:"timestamp"`
}

func NewUserInfo(userid string, pwd string) *UserInfo {
	strtime := FormatCurrentTime()
	userid = strings.ToUpper(userid)
	return &UserInfo{Userid: userid, Pwd: CryptPwd(userid, pwd, strtime), Timestamp: strtime}
}

func (u *UserInfo) GetUserid() string {
	return u.Userid
}

func (u *UserInfo) GetPwd() string {
	return u.Pwd
}

func (u *UserInfo) GetTimestemp() string {
	return u.Timestamp
}

func (u *UserInfo) ParseRecvData(body []byte) bool {
	var r interface{}
	err := json.Unmarshal(body, &r)
	if nil != err {
		fmt.Println("ParseRecvData is error,", err)
	}
	fmt.Println(string(body[:]))
	return true
}
