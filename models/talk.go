package models

import (
	"common/base"
	//"strconv"
	//"errors"
	//"fmt"
	//"strings"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	"github.com/gorilla/websocket"
)

type NewWss struct {
	WsConn   *websocket.Conn
	Ip       string
	UserName string
	Type     string
}

var (
	Wss       = make([]*NewWss, 0)
	Broadcast = make(chan *MessageData)
)

type MessageData struct {
	Message *Talk
}
type Talk struct {
	Id         int64  `orm:"id" json:"id"`
	Fsuname    string `orm:"fsuname" json:"fsuname"`
	Jsuname    string `orm:"jsuname" json:"jsuname"`
	Context    string `orm:"context" json:"context"`
	Newmessage string `orm:"newmessage" json:"newmessage"`
	Ctime      string `orm:"ctime" json:"ctime"`
}

//查询对话参数
type QueryTalkOptions struct {
	BaseOption *base.QueryOptions
	Id         int64
	Fsuname    string
	Jsuname    string
	QingQiuRen string
	Context    string
	Newmessage string
	Ctime      string
}

func init() {
	orm.RegisterModel(new(Talk))
}

//查找对话信息
func QueryTalkInfo(opt *QueryTalkOptions) (int, []*Talk, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Talk))

	cond := orm.NewCondition()
	cond1 := cond.And("fsuname", opt.Fsuname).And("jsuname", opt.Jsuname)
	cond2 := cond.And("fsuname", opt.Jsuname).And("jsuname", opt.Fsuname)
	qs.SetCond(cond1)
	cond3 := cond.AndCond(cond1).OrCond(cond2)
	qs = qs.SetCond(cond3)
	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )

	//	if opt.Fsuname != "" {
	//		qs.Filter("fsuname", opt.Fsuname).Filter("jsuname", opt.Jsuname)
	//	} else {
	//		return 0, nil, nil
	//	}
	//	if opt.Jsuname != "" {
	//		qs.Filter("jsuname__in", opt.Jsuname, opt.Fsuname)
	//	} else {
	//		return 0, nil, nil
	//	}
	var err error
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	talks := make([]*Talk, 0)
	_, err = qs.Limit(1000).OrderBy("ctime").All(&talks)
	if err != nil {
		return 0, nil, err
	}
	if (opt.QingQiuRen == opt.Jsuname) && opt.QingQiuRen != "" {
		_, err = o.QueryTable(new(Talk)).Filter("jsuname", opt.Jsuname).Filter("fsuname", opt.Fsuname).Update(orm.Params{"newmessage": "false"})
	}

	return int(num), talks, err
}

//获取是否有新消息
func QueryNewMessageStatus(opt *QueryTalkOptions) ([]*Talk, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Talk))

	if opt.Jsuname != "" {
		qs = qs.Filter("jsuname", opt.Jsuname)
	} else {
		return nil, nil
	}
	talks := make([]*Talk, 0)
	_, err := qs.Filter("newmessage", "true").All(&talks)
	return talks, err
}

//添加消息
func AddMessage(t *Talk) error {
	o := orm.NewOrm()
	t.Ctime = base.GetCurrentData()
	t.Newmessage = "true"
	_, err := o.Insert(t)
	return err
}

//获取新消息列表
func QueryNewMessageInfo(opt *QueryTalkOptions) ([]*Talk, error) {
	o := orm.NewOrm()
	//qs := o.QueryTable(new(Talk))
	//qs = qs.Filter("ctime__gt", opt.Ctime)

	sql := "SELECT * FROM talk WHERE talk.ctime>'" + opt.Ctime + "' AND talk.fsuname=? AND talk.jsuname=? OR (talk.fsuname=? AND talk.jsuname=? AND talk.ctime>'" + opt.Ctime + "' ) limit 0,1000"
	//	cond := orm.NewCondition()
	//	cond1 := cond.And("fsuname", opt.Fsuname).And("jsuname", opt.Jsuname)
	//	cond2 := cond.And("fsuname", opt.Jsuname).And("jsuname", opt.Fsuname)
	//	cond4 := cond.And("ctime__gt", opt.Ctime)
	//	//	qs.SetCond(cond1)
	//	cond3 := cond.AndCond(cond1).OrCond(cond2).AndCond(cond4)
	//	qs = qs.SetCond(cond3)
	talks := make([]*Talk, 0)
	_, err := o.Raw(sql, opt.Fsuname, opt.Jsuname, opt.Jsuname, opt.Fsuname).QueryRows(&talks)

	//_, err := qs.Limit(1000).All(&talks)
	if opt.QingQiuRen == opt.Jsuname && opt.QingQiuRen != "" {
		_, err := o.QueryTable(new(Talk)).Filter("jsuname", opt.Jsuname).Filter("fsuname", opt.Fsuname).Update(orm.Params{"newmessage": "false"})
		if err != nil {
			return nil, nil
		}
	}

	return talks, err
}
