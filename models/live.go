package models

import (
	"common/base"
	//"errors"
	//"fmt"
	//"strings"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
)

type Live struct {
	Id       int64  `orm:"id"`
	Userid   int64  `orm:"userid"`
	Code     string `orm:"code"`
	Pushflow string `orm:"pushflow"`
	Pullflow string `orm:"pullflow"`
	Ctime    string `orm:"ctime"`
	Info     string `orm:"info"`
	Label    string `orm:"label"`
	Nickname string `orm:"nickname"`
}
type QueryLiveOptions struct {
	base.QueryOptions
	Id     int64
	Userid int64
	Code   string
}

func init() {
	orm.RegisterModel(new(Live))
}
func QueryLiveInfo(opt *QueryLiveOptions) (int, []*Live, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Live))
	cond := orm.NewCondition()

	if opt.Id != 0 {
		cond = cond.And("id", opt.Id)
	}
	if opt.Userid != "" {
		cond = cond.And("userid", opt.Userid)
	}
	if opt.Code != "" {
		cond = cond.And("code", opt.Code)
	}
	qs = qs.SetCond(cond)
	qs.Limit(opt.Limit).Offset(opt.Offset).All()
}
