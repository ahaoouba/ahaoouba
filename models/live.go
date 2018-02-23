package models

import (
	"common/base"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Live struct {
	Id           int64  `orm:"id"`
	Userid       int64  `orm:"userid"`
	Code         string `orm:"code"`
	Pushflow     string `orm:"pushflow"`
	Pullflow     string `orm:"pullflow"`
	Ctime        string `orm:"ctime"`
	Info         string `orm:"info"`
	Label        string `orm:"label"`
	Nickname     string `orm:"nickname"`
	Islive       string `orm:"islive"`
	Lastlinetime string `orm:"lastlinetime"`
}
type QueryLiveOptions struct {
	*base.QueryOptions
	Id     int64
	Userid int64
	Code   string
	Islive string
}

func init() {
	orm.RegisterModel(new(Live))
}

//查询主播信息
func QueryLiveInfo(opt *QueryLiveOptions) (int, []*Live, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Live))
	cond := orm.NewCondition()

	if opt.Id != 0 {
		cond = cond.And("id", opt.Id)
	}
	if opt.Userid != 0 {
		cond = cond.And("userid", opt.Userid)
	}
	if opt.Code != "" {
		cond = cond.And("code", opt.Code)
	}
	if opt.Islive != "" {
		cond = cond.And("islive", opt.Islive)
	}
	qs = qs.SetCond(cond)
	l := make([]*Live, 0)
	num, err := qs.Count()
	if err != nil {
		return 0, nil, err
	}
	_, err = qs.Limit(opt.Limit).Offset(opt.Offset).All(&l)
	return int(num), l, err
}

//添加直播人
func AddLiveinfo(l *Live) error {
	o := orm.NewOrm()
	l.Ctime = base.GetCurrentData()
	l.Code = base.GetUUID()
	err := l.Valited()
	if err != nil {
		return err
	}
	if o.QueryTable(new(Live)).Filter("userid", l.Userid).Exist() {
		return errors.New("sorry,该账号已开通直播间！")
	}
	_, err = o.Insert(l)
	return err
}

//验证添加信息
func (l *Live) Valited() error {
	valid := validation.Validation{}
	valid.Required(l.Userid, "userid")
	valid.Required(l.Code, "code")
	valid.Required(l.Ctime, "ctime")
	valid.Required(l.Nickname, "nickname")

	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

//修改直播信息
func UpdateLiveInfo(l *Live) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Live)).Filter("userid", l.Userid).Update(orm.Params{
		"islive": l.Islive,
	})
	return err
}
