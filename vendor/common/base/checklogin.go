package base

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

func CheckLogin(this beego.Controller) {

	ibyt := make([]byte, 0)
	nbyt := make([]byte, 0)
	abyt := make([]byte, 0)
	ibyt, err := json.Marshal(this.GetSession("userid"))
	if err != nil {
		beego.Error(err)
		return
	}

	nbyt, err = json.Marshal(this.GetSession("username"))
	if err != nil {
		beego.Error(err)
		return
	}
	abyt, err = json.Marshal(this.GetSession("auth"))
	if err != nil {
		beego.Error(err)
		return
	}
	var userid int64
	var username string
	var auth string
	err = json.Unmarshal(ibyt, &userid)
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(nbyt, &username)
	if err != nil {
		beego.Error(err)
		return
	}
	err = json.Unmarshal(abyt, &auth)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["userid"] = userid
	this.Data["username"] = username
	if userid == 0 {
		this.Redirect("/index/login", 302)
		return
	}
	if auth == "user" {
		this.Redirect("/index/index", 302)
		return
	}
}
