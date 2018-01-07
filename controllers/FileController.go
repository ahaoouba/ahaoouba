package controllers

import (
	"ahaoouba/models"
	"strings"
	//	"strconv"
	//	"strings"
	//	"encoding/json"
	"common/ajax"
	"common/base"
	"fmt"
	"os"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

//文件添加页面
func (this *FileController) AddFilePage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/fileadd.html"

}

//文件上传ajax
func (this *FileController) AddFile() {
	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	fil := new(models.File)
	f, fh, err := this.GetFile("file[]")
	if err != nil {
		ar.SetError(fmt.Sprintf("获取文件发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	defer f.Close()
	fil.Fname = fh.Filename

	fil.Filepath = "static" + string(os.PathSeparator) + "file" + string(os.PathSeparator) + fil.Fname
	err = this.SaveToFile("file[]", fil.Filepath)
	if err != nil {
		ar.SetError(fmt.Sprintf("文件保存发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	fil.Ctime = base.GetCurrentData()
	err = models.AddFileInfo(fil)

	if err != nil {
		ar.SetError(fmt.Sprintf("文件信息添加发生异常，错误原因为:[ %s ]!", err))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	ar.Success = true

	this.ServeJSON()
}

//获取文件列表
func (this *FileController) GetFileList() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/filelist.html"
	opt := new(models.QueryFileOptions)
	opt.BaseOption = new(base.QueryOptions)
	var err error
	opt.Id, err = this.GetInt64("id", 0)
	if err != nil {
		beego.Error(err)
		return
	}
	opt.Fname = this.GetString("fname", "")
	opt.BaseOption.Limit = 10
	page, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}
	opt.BaseOption.Offset = (page - 1) * opt.BaseOption.Limit
	num, files, err := models.GetFileInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}
	for _, v := range files {
		v.Filepath = strings.Replace(v.Filepath, "\\", "/", -1)
		v.Filepath = "/" + v.Filepath
	}
	this.Data["files"] = files
	this.Data["page"] = models.NewPage(num, page, 10, this.Ctx.Request.URL.String())
}

//删除文件
func (this *FileController) DelFile() {
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
	err := models.DelFile(id)
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
