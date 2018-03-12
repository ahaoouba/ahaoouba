package models

import (
	"common/base"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Gift struct {
	Id      int64  `orm:"id"`
	Name    string `orm:"name"`
	Picpath string `orm:"picpath"`
	Price   int64  `orm:"price"`
}

type QueryGiftOptions struct {
	*base.QueryOptions
	Id      int64
	Name    string
	Picpath string
	Price   int64
}

func init() {
	orm.RegisterModel(new(Gift))
}

//获取礼物信息
func QueryGiftInfo(opt *QueryGiftOptions) (int, []*Gift, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Gift))
	cond := orm.NewCondition()

	if opt.Id != 0 {
		cond = cond.And("id", opt.Id)
	}
	if opt.Name != "" {
		cond = cond.And("name", opt.Name)
	}
	if opt.Picpath != "" {
		cond = cond.And("picpath", opt.Picpath)
	}
	if opt.Price != 0 {
		cond = cond.And("price", opt.Price)
	}
	qs = qs.SetCond(cond)
	g := make([]*Gift, 0)
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	_, err = qs.Limit(opt.Limit).Offset(opt.Offset).All(&g)
	return int(num), g, err
}

//验证添加信息
func (g *Gift) Valited() error {
	valid := validation.Validation{}
	valid.Required(g.Name, "name")
	valid.Required(g.Picpath, "picpath")
	valid.Required(g.Price, "price")
	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

//添加礼物
func AddGiftInfo(g *Gift) error {
	o := orm.NewOrm()
	err := g.Valited()
	if err != nil {
		return err
	}
	if o.QueryTable(new(Gift)).Filter("name", g.Name).Exist() {
		return errors.New("sorry,该礼物已经存在！")
	}
	_, err = o.Insert(g)
	return err
}

//删除礼物
func DeleteGiftInfo(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Gift)).Filter("id", id).Delete()
	return err
}
