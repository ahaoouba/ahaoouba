package models

import (
	"common/base"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type CateGory struct {
	Id   int64  `orm:"id" json:"id"`
	Pid  int64  `orm:"pid" json:"pid"`
	Name string `orm:"name" json:"name"`
	Jb   int64  `orm:"jb" json:"jb"`
}

func init() {
	orm.RegisterModel(new(CateGory))
}

type QueryCateGoryOptions struct {
	BaseOptions *base.QueryOptions
	Id          int64
	Pid         int64
	Name        string
	Jb          int64
}

//类目添加信息验证
func (this *CateGory) Valited() error {
	valid := validation.Validation{}
	valid.Required(this.Pid, "pid")
	valid.Required(this.Name, "context")
	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

//添加类目
func CateGoryAdd(c *CateGory) error {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CateGory))
	if c.Pid != -1 {
		if !qs.Filter("id", c.Pid).Exist() {
			return errors.New("父级类目不存在!")
		}
	}
	if qs.Filter("name", c.Name).Exist() {
		return errors.New("分类名称已存在!")
	}
	_, err := o.Insert(c)
	return err
}

//类目查询
func QueryCateGoryInfo(opt *QueryCateGoryOptions) (int64, []*CateGory, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CateGory))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Name != "" {
		qs = qs.Filter("name", opt.Name)
	}
	if opt.Pid != 0 {
		qs = qs.Filter("pid", opt.Pid)

	}
	c := make([]*CateGory, 0)
	num, err := qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&c)

	return num, c, err
}

//批量删除类目
func DelCates(ids []int64) error {

	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}
	for _, v := range ids {
		_, err = o.QueryTable(new(CateGory)).Filter("id", v).Delete()
		if err != nil {
			o.Rollback()
			return err
		}
		_, err = o.QueryTable(new(Article)).Filter("cid", v).Delete()
		if err != nil {
			o.Rollback()
			return err
		}
	}
	err = o.Commit()
	return err
}
