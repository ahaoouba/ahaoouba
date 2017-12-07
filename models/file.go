package models

import (
	"common/base"
	"errors"
	//"fmt"
	//"strings"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type File struct {
	Id       int64  `orm:"id"`
	Fname    string `orm:"fname"`
	Filepath string `orm:"filepath"`
	Ctime    string `orm:"ctime"`
}
type QueryFileOptions struct {
	BaseOption *base.QueryOptions
	Id         int64
	Fname      string
	Filepath   string
	Ctime      string
}

func init() {
	orm.RegisterModel(new(File))
}

//获取文件列表
func GetFileInfo(opt *QueryFileOptions) (int, []*File, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(File))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Fname != "" {
		qs = qs.Filter("fname", opt.Fname)
	}
	if opt.Filepath != "" {
		qs = qs.Filter("filepath", opt.Filepath)
	}
	files := make([]*File, 0)
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	_, err = qs.Limit(opt.BaseOption.Limit).Offset(opt.BaseOption.Offset).OrderBy("ctime").All(&files)
	return int(num), files, err
}

//添加文件
func AddFileInfo(f *File) error {
	o := orm.NewOrm()
	if o.QueryTable(new(File)).Filter("fname", f.Fname).Exist() {
		return errors.New("该文件已经存在!")
	}
	_, err := o.Insert(f)
	return err
}

//删除文件
func DelFile(id int64) error {
	o := orm.NewOrm()
	qs := o.QueryTable(new(File)).Filter("id", id)
	if !qs.Exist() {
		return errors.New("删除的目标文件不存在!")
	}
	_, err := qs.Delete()
	return err
}
