package controllers

import (
	"ahaoouba/models"
	"encoding/json"
	"net/http"
	//"strconv"
	//"strings"
	//"encoding/json"
	"common/ajax"
	"common/base"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type TalkController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

//获取对话列表
func (this *TalkController) GetTalkList() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	opt := new(models.QueryTalkOptions)
	opt.BaseOption = new(base.QueryOptions)
	opt.QingQiuRen = this.GetString("qingqiuren", "")
	opt.Fsuname = this.GetString("fsuname", "")
	opt.Jsuname = this.GetString("jsuname", "")
	_, talks, err := models.QueryTalkInfo(opt)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取对话发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Data = talks
	ar.Success = true
	ar.Msg = "获取对话成功!"
	this.ServeJSON()
}

//添加对话信息
func (this *TalkController) AddTalkInfo() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	t := new(models.Talk)
	t.Fsuname = this.GetString("fsuname", "")
	t.Jsuname = this.GetString("jsuname", "")
	t.Context = this.GetString("context", "")
	if t.Fsuname == "" {
		ar.SetError(fmt.Sprintf("发送消息失败，发送者不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if t.Jsuname == "" {
		ar.SetError(fmt.Sprintf("发送消息失败，接收者不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if t.Context == "" {
		ar.SetError(fmt.Sprintf("发送消息失败，发送内容不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err := models.AddMessage(t)
	if err != nil {
		ar.SetError(fmt.Sprintf("发送消息失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	m := new(models.MessageData)
	m.Message = t
	models.Broadcast <- m
	ar.Success = true
	ar.Msg = "发送消息成功!"
	this.ServeJSON()
}

//查看时否有新消息
func (this *TalkController) IsHaveNewmessage() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	opt := new(models.QueryTalkOptions)
	opt.BaseOption = new(base.QueryOptions)
	opt.Jsuname = this.GetString("jsuname", "")
	talks, err := models.QueryNewMessageStatus(opt)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取新消息失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	ntalks := make(map[string]int)
	for _, v := range talks {
		IsNew := true
		if _, exist := ntalks[v.Fsuname]; exist {
			ntalks[v.Fsuname]++
			IsNew = false
		}
		if IsNew {
			ntalks[v.Fsuname] = 1
		}
	}

	ar.Data = ntalks
	ar.Success = true
	ar.Msg = "查看新消息成功!"
	this.ServeJSON()
}

//获取新消息
func (this *TalkController) GetNewMessage() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	opt := new(models.QueryTalkOptions)
	opt.BaseOption = new(base.QueryOptions)
	opt.Jsuname = this.GetString("jsuname", "")
	opt.Fsuname = this.GetString("fsuname", "")
	opt.Ctime = this.GetString("ctime", "")
	opt.QingQiuRen = this.GetString("qingqiuren", "")
	talks, err := models.QueryNewMessageInfo(opt)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取新消息失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Data = talks
	ar.Success = true
	ar.Msg = "新消息获取成功!"
	this.ServeJSON()
}

//websoket请求链接
func (this *TalkController) Ws() {

	upgrader.CheckOrigin = func(Request *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
		return
	}
	wss := new(models.NewWss)
	wss.WsConn = ws
	wss.Ip = this.Ctx.Input.IP()
	wss.Type = this.GetString("type")
	ibyt := make([]byte, 0)
	ibyt, err = json.Marshal(this.GetSession("username"))
	if err != nil {
		beego.Error(err)
		return
	}
	var username string
	err = json.Unmarshal(ibyt, &username)
	if err != nil {
		beego.Error(err)
		return
	}
	wss.UserName = username
	for k, v := range models.Wss {
		if v.Ip == wss.Ip && v.UserName == wss.UserName {
			v.WsConn.Close()
			models.Wss = append(models.Wss[:k], models.Wss[k+1:]...)
		}
	}
	models.Wss = append(models.Wss, wss)
	defer this.ServeJSON()
}
