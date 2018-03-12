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
	///////////////
	/////添加登录状态检测定时任务
	checklogintask := toolbox.NewTask("checklogintasks", "0/10 * * * * *", controllers.CheckLoginStatus)
	toolbox.AddTask("checklogintasks", checklogintask)
	toolbox.StartTask()
	//////////////
	beego.Run("0.0.0.0:8800")
}
