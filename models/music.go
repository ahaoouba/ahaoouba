package models

import (
	"common/base"
	"errors"
	//"errors"
	//"fmt"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Music struct {
	Id    int64  `orm:"id" json:"id"`
	Name  string `orm:"name" json:"name"`
	Path  string `orm:"path" json:"path"`
	Scene string `orm:"scene" json:"scene"`
}

func init() {
	orm.RegisterModel(new(Music))
}

type QueryMusicOptions struct {
	BaseOptions *base.QueryOptions
	Id          int64
	Name        string
	Path        string
	Scene       string
}

func AddMusic(m *Music) error {
	o := orm.NewOrm()
	if o.QueryTable(new(Music)).Filter("name", m.Name).Exist() {
		return errors.New("该文件已经存在!")
	}
	_, err := o.Insert(m)
	return err
}
func QueryMusicInfo(opt *QueryMusicOptions) (int, []*Music, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Music))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Name != "" {
		qs = qs.Filter("name", opt.Name)
	}
	if opt.Path != "" {
		qs = qs.Filter("path", opt.Path)
	}
	if opt.Scene != "" {
		qs = qs.Filter("scene", opt.Scene)
	}
	m := make([]*Music, 0)
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	_, err = qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&m)
	if err != nil {
		return 0, nil, err
	}
	return int(num), m, err
}
func DelMusic(mid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Music)).Filter("id", mid).Delete()
	return err
}
