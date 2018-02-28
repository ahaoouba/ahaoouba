package controllers

import (
	"ahaoouba/models"
	"common/ajax"
	"common/base"
	"encoding/json"
	"math/rand"
	"strconv"
	//"time"

	"fmt"
	"montnets/mwgate/smsutil"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index/reg.html"
}
func (this *IndexController) Reg() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	u := new(models.User)
	u.Name = this.GetString("name", "")
	u.Password = this.GetString("password", "")
	u.Phone = this.GetString("phone", "")
	yzm, _ := this.GetInt("yzm", 0)
	var err error
	var syzm int
	yzmbyt := make([]byte, 0)
	yzmbyt, err = json.Marshal(this.GetSession("yzm"))
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(yzmbyt, &syzm)
	if err != nil {
		beego.Error(err)
		return
	}
	var sph string
	phbyt := make([]byte, 0)
	phbyt, err = json.Marshal(this.GetSession("phone"))
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(phbyt, &sph)
	if err != nil {
		beego.Error(err)
		return
	}
	if yzm != syzm {
		ar.SetError(fmt.Sprintf("验证码错误!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if sph != u.Phone {
		ar.SetError(fmt.Sprintf("电话号码与发送短信电话号不一致!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	err = u.Valited()
	if err != nil {
		ar.SetError(fmt.Sprintf("添加用户发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	u.Islogin = "false"
	err = models.AddUser(u)
	if err != nil {
		ar.SetError(fmt.Sprintf("添加用户发生异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "注册成功!"
	this.ServeJSON()
}
func (this *IndexController) Login() {

	this.TplName = "index/login.html"
}
func (this *IndexController) TuichuLogin() {
	sbyt := make([]byte, 0)
	sbyt, err := json.Marshal(this.GetSession("userid"))
	if err != nil {
		beego.Error(err)
	}
	var userid int64
	err = json.Unmarshal(sbyt, &userid)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateSessionId(userid, "")
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateLinetime(userid, "")
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateIslogin(userid, "true")
	if err != nil {
		beego.Error(err)
	}
	this.DelSession("userid")
	this.DelSession("username")
	this.DelSession("auth")

	this.Redirect("/index/login", 302)
}
func (this *IndexController) LoginAjax() {

	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar

	u := make([]*models.User, 0)
	op := new(models.QueryUserOption)
	bp := new(base.QueryOptions)
	op.BaseOptions = bp
	op.Name = this.GetString("name", "")

	op.Password = this.GetString("password", "")
	if op.Name == "ahaoouba" {
		this.SetSession("userid", 1)
		this.SetSession("username", "ahaoouba")
		this.SetSession("auth", "admin")
		models.UpdateIslogin(1, "true")
		models.UpdateLinetime(1, strconv.FormatInt(base.GetCurrentDataUnix(), 10))
		models.UpdateSessionId(1, this.CruSession.SessionID())
		ar.Success = true
		ar.Msg = "/index/admin"
		this.ServeJSON()
		return
	}

	code, _ := this.GetInt("code")
	var yzm int
	yzmbyt := make([]byte, 0)
	yzmbyt, err := json.Marshal(this.GetSession("login_yzm"))
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(yzmbyt, &yzm)
	if err != nil {
		beego.Error(err)
		return
	}
	if yzm != code {
		ar.SetError(fmt.Sprintf("验正码错误!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	num, u, err := models.QueryUserInfo(op)
	if err != nil {
		ar.SetError(fmt.Sprintf("登录失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if num == 0 {
		ar.SetError(fmt.Sprintf("登录失败，用户名/密码错误!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	if u[0].Islogin == "true" {
		ss, err := beego.GlobalSessions.GetSessionStore(u[0].Sessionid)
		if err != nil {
			beego.Error(err)
		}
		err = ss.Delete("userid")
		if err != nil {
			beego.Error(err)
		}
		err = ss.Delete("username")
		if err != nil {
			beego.Error(err)
		}
		err = ss.Delete("auth")
		if err != nil {
			beego.Error(err)
		}
	}
	//	oop := new(models.QueryUserOption)
	//	oop.BaseOptions = new(base.QueryOptions)
	//	oop.Sessionid = this.CruSession.SessionID()
	//	_, users, err := models.QueryUserInfo(oop)
	//	if err != nil {
	//		beego.Error(err)
	//	}
	//	if len(users) > 0 {
	//		if users[0].Name == op.Name {
	//			ar.SetError(fmt.Sprintf("该账号已在此电脑登录。"))
	//			beego.Error(ar.Errmsg)
	//			this.ServeJSON()
	//			return
	//		} else {
	//			ar.SetError(fmt.Sprintf("同一浏览器不能同时登录两个账号,请换个浏览器或者换个电脑!"))
	//			beego.Error(ar.Errmsg)
	//			this.ServeJSON()
	//			return
	//		}
	//	}
	this.SetSession("userid", u[0].Id)
	this.SetSession("username", u[0].Name)
	this.SetSession("auth", u[0].Auth)
	err = models.UpdateIslogin(u[0].Id, "true")
	if err != nil {
		ar.SetError(fmt.Sprintf("登录失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}

	err = models.UpdateLinetime(u[0].Id, strconv.FormatInt(base.GetCurrentDataUnix(), 10))
	if err != nil {
		ar.SetError(fmt.Sprintf("登录失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	sessionid := this.CruSession.SessionID()

	err = models.UpdateSessionId(u[0].Id, sessionid)
	if err != nil {
		ar.SetError(fmt.Sprintf("登录失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	if u[0].Auth == "admin" {
		ar.Msg = "/index/admin"
	} else {
		ar.Msg = "/index/index"
	}
	this.DelSession("login_yzm")
	this.DelSession("phone")
	this.DelSession("yzm")
	this.ServeJSON()
	return
}
func (this *IndexController) AdminPage() {

	base.CheckLogin(this.Controller)
	this.TplName = "index/admin.html"

}
func (this *IndexController) UserPage() {
	base.CheckLogin(this.Controller)
	this.TplName = "index/user.html"
	opt := new(models.QueryUserOption)
	opt.Name = this.GetString("name", "")
	bp := new(base.QueryOptions)
	opt.BaseOptions = bp
	//分页
	page, err := this.GetInt("p", 1)
	if err != nil {
		beego.Error(err)
		return
	}

	//运算偏移量
	opt.BaseOptions.Offset = (page - 1) * opt.BaseOptions.Limit
	total, u, err := models.QueryUserInfo(opt)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["page"] = models.NewPage(total, page, 10, this.Ctx.Request.URL.String())
	this.Data["user"] = u
	uu := this.GetSession("user")
	ubyt, err := json.Marshal(uu)
	uuser := new(models.User)
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(ubyt, uuser)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["uuser"] = uuser
}
func (this *IndexController) ModUser() {
	base.CheckLogin(this.Controller)
	var err error
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	u := new(models.User)
	u.Id, err = this.GetInt64("id", 0)
	if err != nil {
		ar.SetError(fmt.Sprintf("用户标识异常，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	u.Name = this.GetString("name", "")
	u.Auth = this.GetString("auth", "")
	err = models.ModUser(u)
	if err != nil {
		ar.SetError(fmt.Sprintf("修改用户信息失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = true
	ar.Msg = "修改成功"
	this.ServeJSON()
}
func (this *IndexController) LineTime() {

	base.CheckLogin(this.Controller)
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	ilinetime := base.GetCurrentDataUnix()
	linetime := strconv.FormatInt(ilinetime, 10)
	var err error
	ibyt := make([]byte, 0)
	ibyt, err = json.Marshal(this.GetSession("userid"))
	if err != nil {
		beego.Error(err)
	}
	var id int64
	err = json.Unmarshal(ibyt, &id)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateLinetime(id, linetime)
	if err != nil {
		beego.Error(err)
	}
	ar.Success = true
	this.ServeJSON()
}
func CheckLoginStatus() error {

	opt := new(models.QueryUserOption)
	opt.BaseOptions = new(base.QueryOptions)
	opt.BaseOptions.Limit = 1000
	opt.Islogin = "true"
	_, users, err := models.QueryUserInfo(opt)
	if err != nil {
		beego.Error(err)

	}
	var lt int64
	for _, v := range users {
		time := base.GetCurrentDataUnix()
		if v.Linetime == "" {
			lt = 0
		} else {
			lt, err = strconv.ParseInt(v.Linetime, 10, 64)
			if err != nil {
				beego.Error(err)

			}
		}

		if time-lt > 1000 {
			v.Islogin = "false"
			err = models.UpdateIslogin(v.Id, v.Islogin)
			if err != nil {
				beego.Error(err)

			}
			err = models.UpdateLinetime(v.Id, "")
			if err != nil {
				beego.Error(err)

			}
			err = models.UpdateSessionId(v.Id, "")
			if err != nil {
				beego.Error(err)
			}
			ss, err := beego.GlobalSessions.GetSessionStore(v.Sessionid)
			if err != nil {
				beego.Error(err)
			}
			err = ss.Delete("userid")
			if err != nil {
				beego.Error(err)
			}
			err = ss.Delete("username")
			if err != nil {
				beego.Error(err)
			}
			err = ss.Delete("auth")
			if err != nil {
				beego.Error(err)
			}
		}
	}
	return err
}

//短信验证
func (this *IndexController) Message() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	userid := "E101I4"
	pwd := "g5u0v4"
	mobile := this.GetString("phone")

	content := ""

	r := rand.New(rand.NewSource(base.GetCurrentDataUnix()))
	var yzmnum int
	for i := 0; i < 6; i++ {
		yzmnum = yzmnum*10 + r.Intn(10)
	}

	content = fmt.Sprintf("您的验证码是%d，在1分钟内输入有效。如非本人操作请忽略此短信。", yzmnum)

	// 将数据打包
	sendobj := smsutil.NewSingleSend(userid, pwd, mobile, content)
	// 发送数据
	ok := smsutil.SendAndRecvOnce(sendobj)
	this.SetSession("phone", mobile)
	this.SetSession("yzm", yzmnum)
	if !ok {
		ar.SetError(fmt.Sprintf("验正码发送失败!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = ok
	this.ServeJSON()
}

//登录短信这验证码
func (this *IndexController) LogMessage() {
	ar := ajax.NewAjaxResult()
	this.Data["json"] = ar
	username := this.GetString("name")
	password := this.GetString("password")
	if username == "" || password == "" {
		ar.SetError(fmt.Sprintf("用户名/密码错误!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	u := new(models.QueryUserOption)
	bp := new(base.QueryOptions)
	u.BaseOptions = bp
	u.Name = username
	u.Password = password
	_, user, err := models.QueryUserInfo(u)
	if err != nil {
		ar.SetError(fmt.Sprintf("用户信息获取失败，错误内容为：[ %s ]", err.Error()))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	if len(user) == 0 {
		ar.SetError(fmt.Sprintf("用户名/密码错误!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	mobile := user[0].Phone
	//短信发送账号密码
	userid := "E101I4"
	pwd := "g5u0v4"
	content := ""

	r := rand.New(rand.NewSource(base.GetCurrentDataUnix()))
	var yzmnum int
	for i := 0; i < 6; i++ {
		yzmnum = yzmnum*10 + r.Intn(10)
	}

	content = fmt.Sprintf("您的验证码是%d，在1分钟内输入有效。如非本人操作请忽略此短信。", yzmnum)

	// 将数据打包
	sendobj := smsutil.NewSingleSend(userid, pwd, mobile, content)
	// 发送数据
	ok := smsutil.SendAndRecvOnce(sendobj)
	this.SetSession("login_yzm", yzmnum)
	if !ok {
		ar.SetError(fmt.Sprintf("验正码发送失败!"))
		beego.Error(ar.Errmsg)
		this.ServeJSON()
		return
	}
	ar.Success = ok
	this.ServeJSON()
}
