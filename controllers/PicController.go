package controllers

import (
	"ahaoouba/models"
	"common/base"
	"os"
	"strconv"
	"strings"
	//"time"
	"common/ajax"
	"fmt"

	"github.com/astaxie/beego"
)

type PicController struct {
	beego.Controller
}

//上传图片
func (this *PicController) AddPic() {
	base.CheckLogin(this.Controller)
	p := new(models.Pic)
	p.Show = "true"
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	f, _, err := this.GetFile("file[]")
	if err != nil {
		ar.SetError(fmt.Sprintf("获取图片文件发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	defer f.Close()
	p.Url = "." + string(os.PathSeparator) + "static" + string(os.PathSeparator) + "artimg" + string(os.PathSeparator) + strconv.FormatInt(base.GetCurrentDataUnix(), 10) + "." + "jpg"
	err = this.SaveToFile("file[]", p.Url)
	if err != nil {
		ar.SetError(fmt.Sprintf("图片保存发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err = models.AddPic(p)
	if err != nil {
		ar.SetError(fmt.Sprintf("图片信息添加发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	opt := new(models.QueryPicOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	opt.Url = p.Url
	num, pic, err := models.QueryPicInfo(opt)
	if err != nil {
		ar.SetError(fmt.Sprintf("图片信息获取发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if num == 0 {
		ar.SetError(fmt.Sprintf("图片不存在!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = strconv.FormatInt(pic[0].Id, 10)
	this.ServeJSON()

}

//图片显示隐藏
func (this *PicController) PicShow() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	picurls := this.GetString("urls", "")
	status := this.GetString("status", "")
	urlarr := strings.Split(picurls, ";")
	staarr := strings.Split(status, ";")

	for i := 0; i < len(urlarr); i++ {
		pic := new(models.Pic)
		pic.Show = staarr[i]
		pic.Url = "." + urlarr[i]
		pic.Url = strings.Replace(pic.Url, "/", "\\", -1)
		pic.Url = strings.TrimRight(pic.Url, "\\")
		opt := new(models.QueryPicOptions)
		bp := new(base.QueryOptions)
		opt.BaseOptions = bp
		opt.Url = pic.Url
		beego.Debug(pic.Show)
		num, ss, err := models.QueryPicInfo(opt)
		if err != nil {
			ar.SetError(fmt.Sprintf("图片信息获取发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
		if num != 0 {
			pic.Id = ss[0].Id
		}
		err = models.UpdateShow(pic)
		if err != nil {
			ar.SetError(fmt.Sprintf("图片显隐状态修改发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
	}
	ar.Success = true
	ar.Msg = "操作成功!"
	this.ServeJSON()
}

//图片列表
func (this *PicController) PicList() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/pic.html"
	timearr := make([]string, 0)
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
		v.Ctime = strings.Split(v.Ctime, " ")[0]
		if !base.In_Array(timearr, v.Ctime) {
			timearr = append(timearr, v.Ctime)
		}
	}
	this.Data["pics"] = pics
	this.Data["timearr"] = timearr
}

//删除图片
func (this *PicController) DelPic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	ids := this.GetString("ids", "")
	if ids == "" {
		ar.Success = true
		ar.Msg = "删除成功!"
		this.ServeJSON()
		return
	}
	idarr := strings.Split(ids, ";")
	for _, v := range idarr {
		id, _ := strconv.ParseInt(v, 10, 64)
		err := models.DelPic(id)
		if err != nil {
			ar.SetError(fmt.Sprintf("删除图片发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
		err = models.DelLbPic(id)
		if err != nil {
			ar.SetError(fmt.Sprintf("删除图片发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}

		err = models.DelArtPic(strconv.FormatInt(id, 10))
		if err != nil {
			ar.SetError(fmt.Sprintf("删除图片发生异常，错误原因为:[ %s ]!", err))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
	}
	ar.Msg = "删除成功!"
	this.ServeJSON()
}
