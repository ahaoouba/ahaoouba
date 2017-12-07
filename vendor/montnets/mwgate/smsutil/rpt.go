package smsutil

import "encoding/json"
import "fmt"
import "time"
import "net/url"

type GetRpt struct {
	*UserInfo
	Retsize int `json:"retsize"`
}

type Rpt struct {
	// 平台流水编号
	Msgid uint64 `json:"msgid"`
	// 用户自定义流水编号
	Custid string `json:"custid"`
	// 当前条数
	Pknum int `json:"pknum"`
	// 总条数
	Pktotal int `json:"pktotal"`
	// 短信接收手机号
	Mobile string `json:"mobile"`
	// 完整的通道号
	Spno string `json:"spno"`
	// 扩展号
	Exno string `json:"exno"`
	// 状态报告对应的下行发送时间
	Stime string `json:"stime"`
	// 状态报告返回时间
	Rtime string `json:"rtime"`
	// 接收状态,0:成功 非0:失败
	Status int `json:"status"`
	// 状态报告错误代码
	Errcode string `json:"errcode"`
	// 状态报告错误代码的描述
	Errdesc string `json:"errdesc"`
	// 下行时填写的exdata
	Exdata string `json:"exdata"`
}

// 推送上行接口, 接收云通迅平台推送的MO状态报告
type PushRpt struct {
	*UserInfo
	// 推送上行请求命令：必须填MO_REQ
	Cmd string
	// 请求消息流水号：匹配回应请求的短信包，每次网络请求加1
	Seqid string
	// 上行信息
	rpts []Rpt
}

func NewPushRpt(userid string, pwd string, rpts []Rpt) *PushRpt {
	return &PushRpt{NewUserInfo(userid, pwd), "RPT_REQ", GetCustid(), rpts}
}

func (s *PushRpt) GetName() string {
	return "push_rpt"
}

func (p *PushRpt) GetUrlencode() string {
	v := url.Values{}
	v.Set("userid", p.Userid)
	v.Add("pwd", p.Pwd)
	v.Add("timestamp", p.Timestamp)
	v.Add("cmd", p.Cmd)
	v.Add("seqid", p.Seqid)
	// v.Add("mos", string(PkgToJson(p.Mos)[:]))
	return v.Encode() + "&rpts=" + string(PkgToJson(p.rpts)[:])
}

type RptRet struct {
	Result int   `json:"result"`
	Rpts   []Rpt `json:"rpts"`
}

func NewGetRpt(userid string, pwd string, retsize int) *GetRpt {
	return &GetRpt{NewUserInfo(userid, pwd), retsize}
}

func (s *GetRpt) GetName() string {
	return "get_rpt"
}

func (u *GetRpt) ParseRecvData(body []byte) bool {

	fmt.Println(string(body[:]))

	var r RptRet
	err := json.Unmarshal(body, &r)

	if nil != err {
		fmt.Println("ParseRecvData is error,", err)
	}

	if 0 == r.Result && len(r.Rpts) > 0 {
		fmt.Println("获取上行成功！获取到的上行有", len(r.Rpts), "条记录。")

		for _, v := range r.Rpts {
			fmt.Println("状态报告记录: msgid:", v.Msgid, ",custid:", v.Custid, ",pknum:", v.Pknum, ",pktotal:", v.Pktotal, ",mobile:", v.Mobile, ",spno:", v.Spno, ",exno:", v.Exno, ",stime:", v.Stime, ",rtime:", v.Rtime, ",status:", v.Status, ",errcode:", v.Errcode, ",exdata:", v.Exdata, ",errdesc:", v.Errdesc)
		}
	}

	// if 0 == r["result"] && len(r["rpts"]) > 0 {
	// 	fmt.Println("获取状态报告成功！获取到的状态报告有", rpts.size(), "条记录。")

	// 	for i, v := range r["rpts"] {
	// 		fmt.Println("状态报告记录: msgid:", v["msgid"], ",custid:", v["custid"], ",pknum:", v["pknum"], ",pktotal:", v["pktotal"], ",mobile:", v["mobile"], ",spno:", v["spno"], ",exno:", v["exno"], ",stime:", v["stime"], ",rtime:", v["rtime"], ",status:", v["status"], ",errcode:", v["errcode"], ",exdata:", v["exdata"], ",errdesc:", v["errdesc"])
	// 	}
	// }
	fmt.Println(r)
	return true
}

func (g *GetRpt) run() {

	fmt.Println("GetRpt demo begin")
	for {
		ok := SendAndRecvOnce(g)
		if false == ok {
			// 没有上行，延时5秒以上
			time.Sleep(5000)
		}
	}
}
