package controllers

import (
	"ahaoouba/models"
	"strings"

	"common/ajax"
	"common/base"
	"fmt"
	"os"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

func (this *VideoController) VideoAddPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/video.html"
}

//视频上传ajax
func (this *VideoController) VideoAddAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	fil := new(models.Video)
	f, fh, err := this.GetFile("file[]")
	if err != nil {
		ar.SetError(fmt.Sprintf("获取文件发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	defer f.Close()
	fil.Videoname = fh.Filename

	fil.Videopath = "static" + string(os.PathSeparator) + "video" + string(os.PathSeparator) + fil.Videoname
	err = this.SaveToFile("file[]", fil.Videopath)
	if err != nil {
		ar.SetError(fmt.Sprintf("文件保存发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	fil.Ctime = base.GetCurrentData()
	err = models.AddVideoInfo(fil)

	if err != nil {
		ar.SetError(fmt.Sprintf("文件信息添加发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	ar.Success = true

	this.ServeJSON()
}

//获取视频列表
func (this *VideoController) GetVideoList() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/videolist.html"
	opt := new(models.QueryVideoOptions)
	opt.BaseOption = new(base.QueryOptions)
	var err error
	opt.Id, err = this.GetInt64("id", 0)
	if err != nil {
		beego.Error(err)
		return
	}
	opt.Videoname = this.GetString("videoname", "")
	opt.BaseOption.Limit = 10
	page, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}
	opt.BaseOption.Offset = (page - 1) * opt.BaseOption.Limit
	num, videos, err := models.GetVideoInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range videos {
		v.Videopath = strings.Replace(v.Videopath, "\\", "/", -1)
		v.Videopath = "/" + v.Videopath
	}

	this.Data["videos"] = videos
	this.Data["page"] = models.NewPage(num, page, 10, this.Ctx.Request.URL.String())
}

//删除视频
func (this *VideoController) DelVideo() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	id, _ := this.GetInt64("id", 0)
	path := this.GetString("path", "")

	if id == 0 {
		ar.SetError(fmt.Sprintf("文件标识不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if path == "" {
		ar.SetError(fmt.Sprintf("文件所在路径不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err := models.DelVideo(id)
	if err != nil {
		ar.SetError(fmt.Sprintf("文件删除发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	path = strings.Replace(path, "/", "\\", -1)
	path = strings.Trim(path, "\\")
	err = os.Remove(path)
	if err != nil {
		ar.SetError(fmt.Sprintf("文件删除发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "删除成功！"
	this.ServeJSON()
}

//视频播放页面
func (this *VideoController) VideoPlayPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/videoplay.html"
}
