package controllers

import (
	"ahaoouba/models"
	"strconv"
	"strings"
	//"encoding/json"
	"common/ajax"
	"common/base"

	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

//获取文章列表
func (this *ArticleController) QueryArticle() {

	base.CheckLogin(this.Controller)
	this.TplName = "index/article.html"
	opt := new(models.QueryArticleOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	//分页
	page, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}

	//运算偏移量
	opt.BaseOptions.Offset = (page - 1) * opt.BaseOptions.Limit
	num, art, err := models.QueryArticleInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range art {
		popt := new(models.QueryPicOptions)
		pbp := new(base.QueryOptions)
		popt.BaseOptions = pbp
		ppid := strings.Split(v.Pid, ";")

		for _, j := range ppid {
			if j == "" {
				continue
			}
			popt.Id, _ = strconv.ParseInt(j, 10, 64)
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
	this.Data["page"] = models.NewPage(num, page, 10, this.Ctx.Request.URL.String())
	this.Data["article"] = art
}

//添加文章
func (this *ArticleController) AddArticle() {
	this.TplName = "index/articleadd.html"
	base.CheckLogin(this.Controller)
}

//添加文章ajax
func (this *ArticleController) AddArticleAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	art := new(models.Article)
	art.Title = this.GetString("title", "")
	art.Context = this.GetString("content", "")
	art.Pid = this.GetString("pid", "")
	art.Uid, _ = this.GetInt64("uid", 0)
	art.Cid, _ = this.GetInt64("cid", 0)
	art.Jianjie = this.GetString("jianjie", "")
	art.Ctime = base.GetCurrentData()
	err := art.Valited()
	if err != nil {
		ar.SetError(fmt.Sprintf("文章信息参数异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err = models.ArticleAdd(art)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加文章发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "添加文章成功!"
	this.ServeJSON()
}

//文章添加图片
func (this *ArticleController) AddArtPic() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	id, _ := this.GetInt64("id", 0)
	pid := this.GetString("pid", "")
	if id == 0 || pid == "" {
		ar.SetError(fmt.Sprintf("参数异常!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	err := models.AddArtPic(id, pid)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加图片发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "添加成功!"
	this.ServeJSON()
}

//文章详情页
func (this *ArticleController) ArticleXq() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/artxq.html"
	id, _ := this.GetInt64("id", 0)
	opt := new(models.QueryArticleOptions)
	opt.BaseOptions = new(base.QueryOptions)
	opt.Id = id
	num, art, err := models.QueryArticleInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	if num == 0 {
		return

	}
	uopt := new(models.QueryUserOption)
	uopt.BaseOptions = new(base.QueryOptions)
	uopt.Id = art[0].Uid
	_, u, err := models.QueryUserInfo(uopt)
	if err != nil {
		beego.Error(err)
		return
	}
	art[0].Uname = u[0].Name
	this.Data["art"] = art[0]
	picopt := new(models.QueryPicOptions)
	picopt.BaseOptions = new(base.QueryOptions)
	pids := strings.Split(art[0].Pid, ";")
	if pids[0] == "" {
		return
	}
	picid, err := strconv.ParseInt(pids[0], 10, 64)
	if err != nil {
		beego.Error(err)
		return
	}
	picopt.Id = picid
	_, pic, err := models.QueryPicInfo(picopt)
	if err != nil {
		beego.Error(err)
		return
	}
	urll := pic[0].Url
	urll = strings.TrimLeft(urll, ".")

	art[0].PicUrl = strings.Replace(urll, "\\", "/", -1)

}

//删除文章
func (this *ArticleController) DelArticle() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	id, _ := this.GetInt64("id", 0)
	if id == 0 {
		ar.SetError(fmt.Sprintf("文章标识不能为空!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	err := models.DelArt(id)
	if err != nil {
		ar.SetError(fmt.Sprintf("删除文章发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "删除成功!"
	this.ServeJSON()
}
