package models

import (
	"common/base"
	"errors"
	//"fmt"
	//"strings"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Video struct {
	Id        int64  `orm:"id"`
	Videoname string `orm:"videoname"`
	Videopath string `orm:"videopath"`
	Ctime     string `orm:"ctime"`
}
type QueryVideoOptions struct {
	BaseOption *base.QueryOptions
	Id         int64
	Videoname  string
	Videopath  string
	Ctime      string
}

func init() {
	orm.RegisterModel(new(Video))
}

//获取文件列表
func GetVideoInfo(opt *QueryVideoOptions) (int, []*Video, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Video))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Videoname != "" {
		qs = qs.Filter("videoname", opt.Videoname)
	}
	if opt.Videopath != "" {
		qs = qs.Filter("videopath", opt.Videopath)
	}
	videos := make([]*Video, 0)
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	_, err = qs.Limit(opt.BaseOption.Limit).Offset(opt.BaseOption.Offset).OrderBy("ctime").All(&videos)
	return int(num), videos, err
}

//添加文件
func AddVideoInfo(f *Video) error {
	o := orm.NewOrm()
	if o.QueryTable(new(Video)).Filter("videoname", f.Videoname).Exist() {
		return errors.New("该文件已经存在!")
	}
	_, err := o.Insert(f)
	return err
}

//删除文件
func DelVideo(id int64) error {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Video)).Filter("id", id)
	if !qs.Exist() {
		return errors.New("删除的目标文件不存在!")
	}
	_, err := qs.Delete()
	return err
}
