package models

import (
	"common/base"
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//用户
type User struct {
	Id        int64  `orm:"id" json:"id"`
	Name      string `orm:"name" json:"name"`
	Phone     string `orm:"phone" json:"phone"`
	Password  string `orm:"password" json:"password"`
	Auth      string `orm:"auth" json:"auth"`
	Linetime  string `orm:"linetime" json:"linetime"`
	Islogin   string `orm:"islogin" json:"islogin"`
	Sessionid string `orm:"sessionid" json:"sessionid"`
}

func init() {
	orm.RegisterModel(new(User))
}

type QueryUserOption struct {
	BaseOptions *base.QueryOptions
	Name        string
	Password    string
	Id          int64
	Islogin     string
	Sessionid   string
}

//注册用户
func AddUser(u *User) error {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	if qs.Filter("Name", u.Name).Exist() {
		return errors.New("用户名已存在!")
	}
	u.Auth = "user"
	_, err := o.Insert(u)
	return err
}

//注册字段验证
func (this *User) Valited() error {

	valid := validation.Validation{}
	valid.Required(this.Name, "name")
	valid.Required(this.Password, "password")
	valid.Required(this.Phone, "phone")
	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

//查询用户
func QueryUserInfo(opt *QueryUserOption) (int, []*User, error) {
	o := orm.NewOrm()
	u := make([]*User, 0)
	qs := o.QueryTable(new(User))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Name != "" {
		qs = qs.Filter("name", opt.Name)
	}
	if opt.Password != "" {
		qs = qs.Filter("password", opt.Password)
	}
	if opt.Islogin != "" {
		qs = qs.Filter("islogin", opt.Islogin)
	}
	if opt.Sessionid != "" {
		qs = qs.Filter("sessionid", opt.Sessionid)
	}
	num, err := qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&u)
	if err != nil {
		return 0, nil, err
	}
	return int(num), u, err
}

//修改用户
func ModUser(u *User) error {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	if !qs.Filter("id", u.Id).Exist() {
		return errors.New("修改的目标用户不存在!")
	}
	if !qs.Filter("name", u.Name).Filter("id", u.Id).Exist() {
		if qs.Filter("name", u.Name).Exist() {
			return errors.New("用户已存在!")
		}
	}
	_, err := o.Update(u, "name", "auth")
	return err
}

//更新用户在线时间
func UpdateLinetime(userid int64, linetime string) error {
	o := orm.NewOrm()
	u := new(User)
	u.Id = userid
	u.Linetime = linetime
	_, err := o.Update(u, "linetime")
	return err
}

//更新用户登录状态
func UpdateIslogin(userid int64, islogin string) error {
	o := orm.NewOrm()
	u := new(User)
	u.Id = userid
	u.Islogin = islogin
	_, err := o.Update(u, "islogin")
	return err
}

//更新用户sessionid
func UpdateSessionId(userid int64, sessionid string) error {
	o := orm.NewOrm()
	u := new(User)
	u.Id = userid
	u.Sessionid = sessionid
	_, err := o.Update(u, "sessionid")
	return err
}
