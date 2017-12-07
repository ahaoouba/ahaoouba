package models

import (
	"common/base"
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//文章
type Article struct {
	Id      int64  `orm:"id" json:"id"`
	Title   string `orm:"title" json:"title"`
	Jianjie string `orm:"jianjie" json:"jianjie"`
	Uid     int64  `orm:"uid" json:"uid"`
	Pid     string `orm:"pid" json:"pid"`
	Context string `orm:"context" json:"context"`
	Ctime   string `orm:"ctime" json:"ctime"`
	Cid     int64  `orm:"cid" json:"cid"`
	Cname   string `orm:"-" json:"cname"`
	Uname   string `orm:"-" json:"uname"`
	PicUrl  string `orm:"-" json:"picurl"`
	Shows   string `orm:"-" json:"shows"`
}

//文章查询参数
type QueryArticleOptions struct {
	BaseOptions *base.QueryOptions
	Id          int64
	Uid         int64
	Title       int64
	Ctime       string
	Show        string
	Cid         int64
}

func init() {
	orm.RegisterModel(new(Article))
}

func QueryArticleInfo(opt *QueryArticleOptions) (int, []*Article, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	if opt.Id != 0 {
		qs = qs.Filter("id", opt.Id)
	}
	if opt.Cid != 0 {
		qs = qs.Filter("cid", opt.Cid)
	}
	if opt.Uid != 0 {
		qs = qs.Filter("uid", opt.Uid)
	}

	art := make([]*Article, 0)
	num, err := qs.Limit(opt.BaseOptions.Limit).Offset(opt.BaseOptions.Offset).All(&art)
	return int(num), art, err
}
func ArticleAdd(art *Article) error {
	o := orm.NewOrm()
	if o.QueryTable(new(Article)).Filter("title", art.Title).Exist() {
		return errors.New("文章标题已存在!")
	}
	_, err := o.Insert(art)
	return err
}

//添加文章信息验证
func (this *Article) Valited() error {

	valid := validation.Validation{}
	valid.Required(this.Title, "title")
	valid.Required(this.Context, "context")
	valid.Required(this.Pid, "pid")
	valid.Required(this.Uid, "uid")
	valid.Required(this.Cid, "cid")
	if valid.HasErrors() {
		errmsg := ""
		for _, err := range valid.Errors {
			errmsg = errmsg + fmt.Sprintf(" %s %s ;", err.Key, err.Error())
		}
		return fmt.Errorf("%s", errmsg)
	}
	return nil
}

//添加文章图片
func AddArtPic(id int64, pid string) error {
	var err error
	o := orm.NewOrm()
	art := make([]*Article, 0)
	qs := o.QueryTable(new(Article))
	if qs.Filter("id", id).Exist() {
		_, err := qs.Filter("id", id).All(&art)
		if err != nil {
			return err
		}

		art[0].Pid = art[0].Pid + ";" + pid
		art[0].Pid = strings.TrimLeft(art[0].Pid, ";")
	}
	_, err = o.Update(art[0], "pid")
	return err
}

//删除文章图片
func DelArtPic(pid string) error {
	o := orm.NewOrm()
	arts := make([]*Article, 0)
	num, err := o.QueryTable(new(Article)).Filter("pid__icontains", pid).All(&arts)
	if err != nil {
		return err
	}
	if num == 0 {
		return nil
	}
	for _, v := range arts {
		if strings.Contains(v.Pid, ";") {
			v.Pid = strings.Replace(v.Pid, pid+";", "", -1)
			v.Pid = strings.Replace(v.Pid, pid, "", -1)
		} else {
			v.Pid = strings.Replace(v.Pid, pid, "", -1)
		}

		_, err := o.Update(v, "pid")
		if err != nil {
			return err
		}
	}
	return nil
}

//删除文章
func DelArt(id int64) error {
	o := orm.NewOrm()

	_, err := o.QueryTable(new(Article)).Filter("id", id).Delete()

	return err
}
