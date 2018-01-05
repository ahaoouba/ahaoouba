package main

import (
	"ahaoouba/controllers"
	"ahaoouba/models"
	_ "ahaoouba/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	models.InitDB()
	//添加登录状态检测定时任务
	checklogintask := toolbox.NewTask("checklogintasks", "0/10 * * * * *", controllers.CheckLoginStatus)
	toolbox.AddTask("checklogintasks", checklogintask)
	toolbox.StartTask()
	beego.SetStaticPath("index/staticsd", "stasstic")
	for k, v := range beego.BConfig.WebConfig.StaticDir {
		beego.Debug(k)
		beego.Debug(v)
	}
	//////
	beego.Run("192.168.10.23:8800")

}
