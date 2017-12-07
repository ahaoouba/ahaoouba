package controllers

import (
	"ahaoouba/models"
	"strings"
	//"strconv"
	//"strings"
	//"encoding/json"
	"common/ajax"
	"common/base"

	"fmt"
	"os"

	"github.com/astaxie/beego"
)

type MusicController struct {
	beego.Controller
}

func (this *MusicController) AddMusicPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/musicadd.html"

}
func (this *MusicController) AddMusic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	m := new(models.Music)
	f, fh, err := this.GetFile("file[]")
	if err != nil {
		ar.SetError(fmt.Sprintf("获取音乐文件发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	defer f.Close()
	m.Name = fh.Filename

	beego.Debug(m.Name)
	m.Path = "static" + string(os.PathSeparator) + "music" + string(os.PathSeparator) + m.Name
	err = this.SaveToFile("file[]", m.Path)
	if err != nil {
		ar.SetError(fmt.Sprintf("音乐保存发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err = models.AddMusic(m)
	if err != nil {
		ar.SetError(fmt.Sprintf("音乐信息添加发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	ar.Success = true

	this.ServeJSON()

}
func (this *MusicController) MusicList() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/music.html"
	opt := new(models.QueryMusicOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	opt.Id, _ = this.GetInt64("id", 0)
	opt.Name = this.GetString("name", "")
	opt.Path = this.GetString("path", "")
	opt.Scene = this.GetString("scene", "")
	//分页
	opt.BaseOptions.Limit = 10
	page, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}

	//运算偏移量
	opt.BaseOptions.Offset = (page - 1) * opt.BaseOptions.Limit
	num, m, err := models.QueryMusicInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range m {
		v.Path = strings.Replace(v.Path, "\\", "/", -1)
	}
	this.Data["page"] = models.NewPage(num, page, 10, this.Ctx.Request.URL.String())
	this.Data["music"] = m

}

//删除音乐
func (this *MusicController) DelMusic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	mid, _ := this.GetInt64("mid", 0)
	path := this.GetString("path", "")
	if mid == 0 {
		ar.SetError(fmt.Sprintf("音乐标识不能为空!"))
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
	err := models.DelMusic(mid)
	if err != nil {
		ar.SetError(fmt.Sprintf("音乐信息删除发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err = os.Remove(path)
	if err != nil {
		ar.SetError(fmt.Sprintf("音乐文件删除发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "删除成功！"
	this.ServeJSON()
}
