package models

import (
	"common/base"
	"errors"
	"os"
	"strings"
	//"errors"
	//"fmt"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Pic struct {
	Id    int64  `orm:"id" json:"id"`
	Url   string `orm:"url" json:"url"`
	Show  string `orm:"show" json:"show"`
	Ctime string `orm:"ctime" json:"ctime"`
}

func init() {
	orm.RegisterModel(new(Pic))
}

type QueryPicOptions struct {
	BaseOptions *base.QueryOptions
	Id          int64
	Show        string
	Url         string
	Ctime       string
}

func AddPic(p *Pic) error {
	o := orm.NewOrm()
	p.Ctime = base.GetCurrentData()
	_, err := o.Insert(p)
	return err
}
func QueryPicInfo(opt *QueryPicOptions) (int64, []*Pic, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pic))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Show != "" {
		qs = qs.Filter("show", opt.Show)
	}
	if opt.Url != "" {
		qs = qs.Filter("url", opt.Url)
	}
	if opt.Ctime != "" {
		qs = qs.Filter("ctime__icontains", opt.Ctime)
	}
	p := make([]*Pic, 0)
	num, err := qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&p)
	return num, p, err
}

//更新图片显示状态
func UpdateShow(pic *Pic) error {
	o := orm.NewOrm()
	_, err := o.Update(pic, "show")
	return err
}

//删除图片
func DelPic(id int64) error {
	o := orm.NewOrm()
	pic := make([]*Pic, 0)
	num, err := o.QueryTable(new(Pic)).Filter("id", id).All(&pic)
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("图片不存在!")
	}
	surl := strings.TrimLeft(pic[0].Url, ".")
	surl = strings.Replace(surl, "\\", "/", -1)
	surl = strings.Trim(surl, "/")
	err = os.Remove(surl)
	if err != nil {
		return err
	}
	_, err = o.QueryTable(new(Pic)).Filter("id", id).Delete()
	return err
}
