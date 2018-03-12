package controllers

import (
	"ahaoouba/models"
	"common/ajax"
	"common/base"
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
)

type GiftController struct {
	beego.Controller
}

//获取礼物列表
func (this *GiftController) QueryGiftInfoAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar

	op := new(models.QueryGiftOptions)
	bp := new(base.QueryOptions)
	op.QueryOptions = bp
	op.Limit = 1000
	var err error
	op.Id, err = this.GetInt64("id", 0)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取礼物ID发生异常：[%s]", err.Error()))
		beego.Error(err)
		ar.Success = false
		this.ServeJSON()
	}
	op.Picpath = this.GetString("picpath", "")
	_, gifts, err := models.QueryGiftInfo(op)
	if err != nil {
		ar.SetError("获取礼物信息发生异常！")
		beego.Error(err)
		ar.Success = false
		this.ServeJSON()
	}
	ar.Data = gifts
	ar.Success = true
	this.ServeJSON()
}

//添加礼物页面
func (this *GiftController) AddGiftPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/addgift.html"
}

//添加礼物
func (this *GiftController) AddGiftAjax() {
	var err error
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	g := new(models.Gift)
	g.Name = this.GetString("name", "")
	g.Picpath = this.GetString("picpath", "")
	g.Price, err = this.GetInt64("price", 0)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取礼物价格发生异常：[%s]", err.Error()))
		beego.Error(err)
		ar.Success = false
		this.ServeJSON()
	}
	err = models.AddGiftInfo(g)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加礼物发生异常：[%s]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "添加礼物成功"
	this.ServeJSON()
}

//添加礼物图片
func (this *GiftController) AddGiftPic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	////礼物图片
	f, fh, err := this.GetFile("file[]")
	if err != nil {
		ar.SetError(fmt.Sprintf("获取文件发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	defer f.Close()
	fname := base.GetUUID() + "." + strings.Split(fh.Filename, ".")[1]
	Picpath := "static" + string(os.PathSeparator) + "gift" + string(os.PathSeparator) + fname
	err = this.SaveToFile("file[]", Picpath)
	if err != nil {
		ar.SetError(fmt.Sprintf("文件保存发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	////
	ar.Data = Picpath
	ar.Success = true
	this.ServeJSON()
}

//删除礼物
func (this *GiftController) DeleteGiftAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	id, err := this.GetInt64("id", 0)
	if err != nil {
		ar.SetError(fmt.Sprintf("获取礼物ID发生异常：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	err = models.DeleteGiftInfo(id)
	if err != nil {
		ar.SetError(fmt.Sprintf("删除礼物发生异常：[%s]", err.Error()))
		beego.Error(err)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "删除成功!"
	this.ServeJSON()
}
