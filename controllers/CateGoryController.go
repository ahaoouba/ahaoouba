package controllers

import (
	"ahaoouba/models"
	"common/ajax"
	"common/base"
	//"encoding/json"

	"fmt"

	"github.com/astaxie/beego"
)

type CateGoryController struct {
	beego.Controller
}

func (this *CateGoryController) AddCateGoryPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/cateadd.html"
}

//分类数据模型
func (this *CateGoryController) CateModel() {

	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	opt := new(models.QueryCateGoryOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	opt.BaseOptions.Limit = 1000
	_, c, err := models.QueryCateGoryInfo(opt)

	if err != nil {
		ar.SetError(fmt.Sprintf("分类信息获取异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	nc := wxfl(c, -1, 1)
	ar.Data = nc
	this.ServeJSON()
	newc = nil
}

var newc = make([]*models.CateGory, 0)

func wxfl(c []*models.CateGory, pid int64, jb int64) []*models.CateGory {
	for _, v := range c {
		if v.Pid == pid {
			v.Jb = jb
			newc = append(newc, v)
			wxfl(c, v.Id, v.Jb+1)
		}
	}
	return newc
}
func (this *CateGoryController) CateAddAjax() {
	base.CheckLogin(this.Controller)
	var err error
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	c := new(models.CateGory)
	c.Pid, _ = this.GetInt64("pid", 0)

	c.Name = this.GetString("name", "")
	if c.Pid == 0 {
		c.Jb = 1
	} else if c.Pid == -1 {
		c.Jb = 2
	} else {
		opt := new(models.QueryCateGoryOptions)
		bp := new(base.QueryOptions)
		opt.BaseOptions = bp
		opt.Id = c.Pid
		_, ca, err := models.QueryCateGoryInfo(opt)
		if err != nil {
			ar.SetError(fmt.Sprintf("分类信息获取异常，错误内容为：[ %s ]", err.Error()))
			beego.Error(ar.Errmsg)
			this.ServeJSON()
			return
		}
		c.Jb = ca[0].Jb + 1
	}
	if c.Pid == 0 {
		c.Pid = -1
	}
	err = models.CateGoryAdd(c)
	if err != nil {
		ar.SetError(fmt.Sprintf("分类信息获取异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "添加类目成功!"
	this.ServeJSON()

}

//删除类目跳转页面
func (this *CateGoryController) DelCatePage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/delcate.html"
}

//删除类目ajax
func (this *CateGoryController) DelCateAjax() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	cid, _ := this.GetInt64("cid", 0)
	if cid == 0 {
		ar.SetError(fmt.Sprintf("类目标识获取异常!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	//查询该类目id下所有子类目
	opt := new(models.QueryCateGoryOptions)
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	opt.BaseOptions.Limit = 1000
	_, c, err := models.QueryCateGoryInfo(opt)

	if err != nil {
		ar.SetError(fmt.Sprintf("分类信息获取异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	nnc = append(nnc, cid)
	nnc = Px(c, cid)
	err = models.DelCates(nnc)
	ar.Msg = "删除成功!"
	if err != nil {
		ar.SetError(fmt.Sprintf("删除类目发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	this.ServeJSON()
	nnc = nil
}

var nnc = make([]int64, 0)

func Px(c []*models.CateGory, cid int64) []int64 {
	for i := 0; i < len(c); i++ {
		if c[i].Pid == cid {
			nnc = append(nnc, c[i].Id)
			Px(c, c[i].Id)
		}
	}
	return nnc
}
