package controllers

import (
	"common/base"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

func (this *VideoController) VideoAddPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/video.html"
}
