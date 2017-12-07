package controllers

import (
	"ahaoouba/models"
	"strconv"
	"strings"
	//	"common/ajax"
	"common/base"
	//	"encoding/json"

	//	"fmt"

	"github.com/astaxie/beego"
)

type QtIndexController struct {
	beego.Controller
}

//前台主页
func (this *QtIndexController) IndexPage() {

	this.TplName = "qtmb/index.html"
	//获取录播图片
	opt := new(models.QueryLbpicOptions)
	opt.BaseOptions = new(base.QueryOptions)
	opt.BaseOptions.Limit = 5
	_, lbpics, err := models.QueryLbpicInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range lbpics {
		v.Urlpath = strings.Replace(v.Urlpath, "\\", "/", -1)
		v.Urlpath = strings.TrimLeft(v.Urlpath, ".")
	}
	this.Data["lbpics"] = lbpics

}

//文章列表页
func (this *QtIndexController) ArticleListPage() {
	this.TplName = "qtmb/blog.html"
	cid, err := this.GetInt64("cid", 0)
	if err != nil {
		beego.Error(err)
		return
	}
	//分页
	p, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}
	opt := new(models.QueryArticleOptions)
	opt.BaseOptions = new(base.QueryOptions)
	opt.BaseOptions.Limit = 10

	opt.BaseOptions.Offset = (p - 1) * opt.BaseOptions.Limit
	opt.Cid = cid
	num, arts, err := models.QueryArticleInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	/////////
	for _, v := range arts {

		popt := new(models.QueryPicOptions)
		pbp := new(base.QueryOptions)
		popt.BaseOptions = pbp
		ppid := strings.Split(v.Pid, ";")

		for _, j := range ppid {
			if j == "" {
				continue
			}
			popt.Id, _ = strconv.ParseInt(j, 10, 64)
			popt.Show = "true"
			pnum, pic, err := models.QueryPicInfo(popt)
			if err != nil {
				beego.Error(err)
				return
			}
			if pnum != 0 {

				str := strings.Replace(pic[0].Url, "\\", "/", -1)
				v.PicUrl = v.PicUrl + ";" + strings.TrimLeft(str, ".")
				v.PicUrl = strings.TrimLeft(v.PicUrl, ";")
				v.Shows = v.Shows + ";" + pic[0].Show
				v.Shows = strings.TrimLeft(v.Shows, ";")
				break

			} else {
				v.PicUrl = v.PicUrl + ""

			}
		}

		uopt := new(models.QueryUserOption)
		ubp := new(base.QueryOptions)
		uopt.BaseOptions = ubp
		uopt.Id = v.Uid
		unum, user, err := models.QueryUserInfo(uopt)

		if err != nil {
			beego.Error(err)
			return
		}
		if unum != 0 {
			v.Uname = user[0].Name
		}
		copt := new(models.QueryCateGoryOptions)
		cbp := new(base.QueryOptions)
		copt.BaseOptions = cbp
		copt.Id = v.Cid
		cnum, cate, err := models.QueryCateGoryInfo(copt)
		if err != nil {
			beego.Error(err)
			return
		}
		if cnum != 0 {
			v.Cname = cate[0].Name
		}

	}
	///////////
	this.Data["arts"] = arts
	this.Data["page"] = models.NewPage(num, p, 10, this.Ctx.Request.URL.String())

}

//文章详情页
func (this *QtIndexController) ArticleXq() {

}
