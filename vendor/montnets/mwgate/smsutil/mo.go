package smsutil

import "encoding/json"
import "fmt"
import "time"
import "net/url"

type Mo struct {
	// 平台流水编号
	Msgid string `json:"msgid"`
	// 上行手机号
	Mobile string `json:"mobile"`
	// 上行通道号
	Spno string `json:"spno"`
	// 上行扩展子号
	Exno string `json:"exno"`
	// 上行时间
	Rtime string `json:"rtime"`
	// 上行内容
	Content string `json:"content"`
}

type MoRet struct {
	Result int  `json:"result"`
	Mos    []Mo `json:"mos"`
}

// 推送上行接口, 接收云通迅平台推送的MO状态报告
type PushMo struct {
	*UserInfo
	// 推送上行请求命令：必须填MO_REQ
	Cmd string
	// 请求消息流水号：匹配回应请求的短信包，每次网络请求加1
	Seqid string
	// 上行信息
	Mos []Mo
}

func NewPushMo(userid string, pwd string, mos []Mo) *PushMo {
	return &PushMo{NewUserInfo(userid, pwd), "MO_REQ", GetCustid(), mos}
}

func (s *PushMo) GetName() string {
	return "push_mo"
}

func (p *PushMo) GetUrlencode() string {
	v := url.Values{}
	v.Set("userid", p.Userid)
	v.Add("pwd", p.Pwd)
	v.Add("timestamp", p.Timestamp)
	v.Add("cmd", p.Cmd)
	v.Add("seqid", p.Seqid)
	// v.Add("mos", string(PkgToJson(p.Mos)[:]))
	return v.Encode() + "&mos=" + string(PkgToJson(p.Mos)[:])
}

// 获取上行接口
type GetMo struct {
	*UserInfo
	Retsize int `json:"retsize"`
}

func NewGetMo(userid string, pwd string, retsize int) *GetMo {
	return &GetMo{NewUserInfo(userid, pwd), retsize}
}

func (s *GetMo) GetName() string {
	return "get_mo"
}

func (u *GetMo) ParseRecvData(body []byte) bool {
	//var r interface{}
	r := &MoRet{Result: 0, Mos: make([]Mo, 0)}
	err := json.Unmarshal(body, &r)

	if nil != err {
		fmt.Println("ParseRecvData is error,", err)
	}

	if 0 == r.Result && len(r.Mos) > 0 {
		fmt.Println("获取上行成功！获取到的上行有", len(r.Mos), "条记录。")

		for _, v := range r.Mos {
			fmt.Println("上行记录: msgid:", v.Msgid, ",mobile:", v.Mobile, ",spno:", v.Spno, ",exno:", v.Exno, ",rtime:", v.Rtime, ",content:", v.Content)
		}
	}
	fmt.Println(r)
	return true
}

func (g *GetMo) run() {

	fmt.Println("getMo demo begin")
	for {
		ok := SendAndRecvOnce(g)
		if false == ok {
			// 没有上行，延时5秒以上
			time.Sleep(5000)
		}
	}
}
