package controllers

import (
	"ahaoouba/models"
	"encoding/json"

	"common/ajax"
	"common/base"

	"github.com/astaxie/beego"
)

const ZeroTime string = "0000-00-00 00:00:00"

type LiveController struct {
	beego.Controller
}

//申请直播间页面
func (this *LiveController) AddLiveinfoPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/liveadd.html"
}

//申请直播间Ajax
func (this *LiveController) AddLiveinfoAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	l := new(models.Live)
	l.Info = this.GetString("info", "")
	l.Nickname = this.GetString("nickname", "")
	l.Label = this.GetString("label", "")
	uidbyt, err := json.Marshal(this.GetSession("userid"))
	if err != nil {
		ar.SetError(err.Error())
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err = json.Unmarshal(uidbyt, &l.Userid)
	if err != nil {
		ar.SetError(err.Error())
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	l.Islive = "false"
	l.Lastlinetime = ZeroTime
	err = models.AddLiveinfo(l)

	if err != nil {
		ar.SetError(err.Error())
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}

//查看用户直播间信息
func (this *LiveController) QueryLiveInfo() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/mylive.html"
	opt := new(models.QueryLiveOptions)
	bp := new(base.QueryOptions)
	opt.QueryOptions = bp
	uidbyt, err := json.Marshal(this.GetSession("userid"))
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(uidbyt, &opt.Userid)
	if err != nil {
		beego.Error(err)
		return
	}
	num, lives, err := models.QueryLiveInfo(opt)
	if num == 0 {
		beego.Error("未申请直播间!")
		return
	}
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["lives"] = lives[0]

}

//修改直播信息
func (this *LiveController) UpdateLiveInfo() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	l := new(models.Live)
	if this.GetString("islive", "") == "true" {
		l.Islive = "true"
	} else if this.GetString("islive", "") == "false" {
		l.Islive = "false"
	} else {
		l.Islive = ""
	}
	uidbyt, err := json.Marshal(this.GetSession("userid"))
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(uidbyt, &l.Userid)
	if err != nil {
		beego.Error(err)
		return
	}
	err = models.UpdateLiveInfo(l)
	if err != nil {
		ar.SetError(err.Error())
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	this.ServeJSON()
}

//直播大厅
func (this *LiveController) LiveHall() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/livehall.html"
	opt := new(models.QueryLiveOptions)
	bp := new(base.QueryOptions)
	opt.QueryOptions = bp
	opt.Islive = "true"
	num, lives, err := models.QueryLiveInfo(opt)
	if num == 0 {
		return
	}
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["lives"] = lives
}

//进入直播间
func (this *LiveController) LivePlayRoom() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/liveplay.html"
	id, err := this.GetInt64("id", 0)
	if err != nil {
		beego.Error(err)
		return
	}
	opt := new(models.QueryLiveOptions)
	bp := new(base.QueryOptions)
	opt.QueryOptions = bp
	opt.Id = id
	_, l, err := models.QueryLiveInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["live"] = l[0]

}
