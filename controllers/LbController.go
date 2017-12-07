package controllers

import (
	"ahaoouba/models"
	"common/ajax"
	"common/base"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type LbController struct {
	beego.Controller
}

func (this *LbController) LbpicPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/lbpic.html"
	opt := new(models.QueryPicOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	opt.BaseOptions.Limit = 1000
	_, pics, err := models.QueryPicInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range pics {
		v.Url = strings.TrimLeft(v.Url, ".")
		v.Url = strings.Replace(v.Url, "\\", "/", -1)
	}
	oopt := new(models.QueryLbpicOptions)
	bbp := new(base.QueryOptions)
	oopt.BaseOptions = bbp
	oopt.BaseOptions.Limit = 1000
	_, lbpics, err := models.QueryLbpicInfo(oopt)
	if err != nil {
		beego.Error(err)
	}
	var oldids = ""
	for _, v := range lbpics {
		oldids = oldids + ";" + strconv.FormatInt(v.Pid, 10)
	}
	oldids = strings.TrimLeft(oldids, ";")
	this.Data["oldids"] = oldids
	this.Data["pics"] = pics
}
func (this *LbController) SzLbPic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	var err error
	this.Data["json"] = ar
	ids := this.GetString("ids", "")
	idarr := strings.Split(ids, ";")
	for _, v := range idarr {
		id, _ := strconv.ParseInt(v, 10, 64)
		err = models.AddLbPic(id)
		if err != nil {
			ar.SetError(fmt.Sprintf("添加轮播图片发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
	}
	delids := this.GetString("delids", "")
	delidarr := strings.Split(delids, ";")
	for _, v := range delidarr {
		pid, _ := strconv.ParseInt(v, 10, 64)
		err = models.DelLbPic(pid)
		if err != nil {
			ar.SetError(fmt.Sprintf("添加轮播图片发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
	}
	ar.Success = true
	ar.Msg = "轮播图片设置成功!"
	this.ServeJSON()
}
