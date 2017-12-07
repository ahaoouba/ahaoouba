package models

import (
	"common/base"

	//"fmt"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Lbpic struct {
	Id      int64  `orm:"id" json:"id"`
	Pid     int64  `orm:"pid" json:"pid"`
	Urlpath string `orm:"urlpath" json:"urlpath"`
}

func init() {
	orm.RegisterModel(new(Lbpic))
}

type QueryLbpicOptions struct {
	BaseOptions *base.QueryOptions
	Id          int64
	Pid         int64
	Urlpath     string
}

func QueryLbpicInfo(opt *QueryLbpicOptions) (int64, []*Lbpic, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lbpic))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Pid != 0 {
		qs = qs.Filter("pid", opt.Pid)
	}
	if opt.Urlpath != "" {
		qs = qs.Filter("urlpath", opt.Urlpath)
	}
	lbpicarr := make([]*Lbpic, 0)
	num, err := qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&lbpicarr)
	if err != nil {
		return 0, nil, err
	}
	return num, lbpicarr, err
}
func AddLbPic(pid int64) error {
	o := orm.NewOrm()
	pic := new(Pic)
	err := o.QueryTable(new(Pic)).Filter("id", pid).One(pic)
	if err != nil {
		return err
	}
	if o.QueryTable(new(Lbpic)).Filter("pid", pid).Exist() {
		return nil
	}
	lbpic := new(Lbpic)
	lbpic.Pid = pic.Id
	lbpic.Urlpath = pic.Url
	_, err = o.Insert(lbpic)
	return err
}

//根据pid删除轮播图片
func DelLbPic(pid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Lbpic)).Filter("pid", pid).Delete()
	return err
}
