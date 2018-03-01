package controllers

import (
	"ahaoouba/models"
	"common/ajax"
	"common/base"

	"net/http"

	"github.com/astaxie/beego"
)

type TalkRoomController struct {
	beego.Controller
}

//聊天室创建连接
func (this *TalkRoomController) TalkRoom() {

	upgrader.CheckOrigin = func(Request *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
		return
	}
	wss := new(models.TalkRoom)
	wss.WsConn = ws
	wss.RoomCode = this.GetString("code", "")
	wss.Ip = this.Ctx.Input.IP()
	for k, v := range models.RoomWss {
		if v.Ip == wss.Ip && v.RoomCode == wss.RoomCode {
			v.WsConn.Close()
			models.RoomWss = append(models.RoomWss[:k], models.RoomWss[k+1:]...)
		}

	}
	models.RoomWss = append(models.RoomWss, wss)
	go func() {
		rt := new(models.RoomTalk)
		m := new(models.RoomTalkData)

		rt.Content = "欢迎" + this.GetString("username") + "来到本直播间!"
		rt.Code = this.GetString("code", "")
		rt.Fstime = base.GetCurrentData()
		rt.Type = models.HuanYing

		m.Message = rt
		models.RoomBroadcast <- m
	}()

	defer this.ServeJSON()
}

//发送对话信息
func (this *TalkRoomController) TalkRoomAddMessage() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	rt := new(models.RoomTalk)
	rt.Content = this.GetString("content", "")
	rt.Fsname = this.GetString("fsname", "")
	rt.Code = this.GetString("code", "")
	rt.Fstime = base.GetCurrentData()
	rt.Type = models.LiaoTian
	m := new(models.RoomTalkData)
	m.Message = rt
	models.RoomBroadcast <- m
	ar.Success = true
	this.ServeJSON()
}
