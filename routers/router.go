package routers

import (
	"ahaoouba/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.IndexController{}, "GET:Login")
	ns := beego.NewNamespace("index/",

		//后台
		////////////////
		///////////////
		//注册页面
		beego.NSRouter("reg", &controllers.IndexController{}),
		//注册ajax
		beego.NSRouter("regajax", &controllers.IndexController{}, "POST:Reg"),
		//注册短信验证
		beego.NSRouter("message", &controllers.IndexController{}, "POST:Message"),
		//登录页面
		beego.NSRouter("login", &controllers.IndexController{}, "GET:Login"),
		beego.NSRouter("tlogin", &controllers.IndexController{}, "GET:TuichuLogin"),
		//登录短信验正码
		beego.NSRouter("logmessage", &controllers.IndexController{}, "POST:LogMessage"),
		//更新用户在线时间
		beego.NSRouter("linetime", &controllers.IndexController{}, "GET:LineTime"),
		//登录ajax
		beego.NSRouter("loginajax", &controllers.IndexController{}, "POST:LoginAjax"),
		//管理员主页
		beego.NSRouter("admin", &controllers.IndexController{}, "GET:AdminPage"),
		//用户管理
		beego.NSRouter("user", &controllers.IndexController{}, "GET:UserPage"),
		beego.NSRouter("moduser", &controllers.IndexController{}, "POST:ModUser"),
		//文章管理
		beego.NSRouter("article", &controllers.ArticleController{}, "GET:QueryArticle"),
		beego.NSRouter("delart", &controllers.ArticleController{}, "POST:DelArticle"),
		beego.NSRouter("articleadd", &controllers.ArticleController{}, "GET:AddArticle"),
		beego.NSRouter("addartpic", &controllers.ArticleController{}, "POST:AddArtPic"),
		beego.NSRouter("artxq", &controllers.ArticleController{}, "GET:ArticleXq"),
		beego.NSRouter("artaddajax", &controllers.ArticleController{}, "POST:AddArticleAjax"),
		//图片管理
		beego.NSRouter("picadd", &controllers.PicController{}, "POST:AddPic"),
		beego.NSRouter("picshow", &controllers.PicController{}, "POST:PicShow"),
		beego.NSRouter("piclist", &controllers.PicController{}, "GET:PicList"),
		beego.NSRouter("delpic", &controllers.PicController{}, "POST:DelPic"),
		//分类管理
		beego.NSRouter("cateAdd", &controllers.CateGoryController{}, "GET:AddCateGoryPage"),
		beego.NSRouter("catemodel", &controllers.CateGoryController{}, "GET:CateModel"),
		beego.NSRouter("delcatepage", &controllers.CateGoryController{}, "GET:DelCatePage"),
		beego.NSRouter("catedelajax", &controllers.CateGoryController{}, "GET:DelCateAjax"),
		beego.NSRouter("cateaddajax", &controllers.CateGoryController{}, "POST:CateAddAjax"),
		//轮播管理
		beego.NSRouter("lbpic", &controllers.LbController{}, "GET:LbpicPage"),
		beego.NSRouter("szlbpic", &controllers.LbController{}, "POST:SzLbPic"),
		//音乐管理
		beego.NSRouter("musicaddpage", &controllers.MusicController{}, "GET:AddMusicPage"),
		beego.NSRouter("musicadd", &controllers.MusicController{}, "POST:AddMusic"),
		beego.NSRouter("music", &controllers.MusicController{}, "GET:MusicList"),
		beego.NSRouter("delmusic", &controllers.MusicController{}, "POST:DelMusic"),
		//视频管理
		beego.NSRouter("videoadd", &controllers.VideoController{}, "GET:VideoAddPage"),
		beego.NSRouter("videoadd", &controllers.VideoController{}, "POST:VideoAddAjax"),
		beego.NSRouter("video", &controllers.VideoController{}, "GET:GetVideoList"),
		beego.NSRouter("videodel", &controllers.VideoController{}, "POST:DelVideo"),
		beego.NSRouter("videoplay", &controllers.VideoController{}, "GET:VideoPlayPage"),
		//文件管理
		beego.NSRouter("fileadd", &controllers.FileController{}, "GET:AddFilePage"),
		beego.NSRouter("fileadd", &controllers.FileController{}, "POST:AddFile"),
		beego.NSRouter("filelist", &controllers.FileController{}, "GET:GetFileList"),
		beego.NSRouter("filedel", &controllers.FileController{}, "POST:DelFile"),
		//消息管理
		//获取所有消息
		beego.NSRouter("talks", &controllers.TalkController{}, "POST:GetTalkList"),
		//添加对话信息
		beego.NSRouter("addtalk", &controllers.TalkController{}, "POST:AddTalkInfo"),
		//查看是否有新消息
		beego.NSRouter("ishavenewtalk", &controllers.TalkController{}, "POST:IsHaveNewmessage"),
		//获取新消息信息
		beego.NSRouter("newtalkinfo", &controllers.TalkController{}, "POST:GetNewMessage"),
		//websoket
		beego.NSRouter("ws", &controllers.TalkController{}, "GET:Ws"),
		///////////
		//直播管理
		//申请直播页面
		beego.NSRouter("liveaddpage", &controllers.LiveController{}, "GET:AddLiveinfoPage"),
		//申请直播ajax
		beego.NSRouter("liveadd", &controllers.LiveController{}, "POST:AddLiveinfoAjax"),
		//用户本人直播间信息
		beego.NSRouter("myliveinfo", &controllers.LiveController{}, "GET:QueryLiveInfo"),
		//修改用户直播状态
		beego.NSRouter("startstoplive", &controllers.LiveController{}, "POST:UpdateLiveInfo"),
		//直播大厅
		beego.NSRouter("livehall", &controllers.LiveController{}, "GET:LiveHall"),
		//进入直播间
		beego.NSRouter("liveplay", &controllers.LiveController{}, "GET:LivePlayRoom"),
		////////////
		//聊天室
		///////////
		beego.NSRouter("wstalkroom", &controllers.TalkRoomController{}, "GET:TalkRoom"),
		beego.NSRouter("talkroommessages", &controllers.TalkRoomController{}, "POST:TalkRoomAddMessage"),
		/////////
		//礼物
		//////////
		//获取礼物信息
		beego.NSRouter("getgifts", &controllers.GiftController{}, "GET:QueryGiftInfoAjax"),
		///添加礼物信息
		beego.NSRouter("addgift", &controllers.GiftController{}, "POST:AddGiftAjax"),
		//添加礼物展示页面
		beego.NSRouter("addgiftpage", &controllers.GiftController{}, "GET:AddGiftPage"),
		//删除礼物
		beego.NSRouter("deletegift", &controllers.GiftController{}, "POST:DeleteGiftAjax"),
		//添加礼物图片
		beego.NSRouter("addgiftpic", &controllers.GiftController{}, "POST:AddGiftPic"),
		/////////////////
		//前台
		////////////////
		///////////////
		//前台主页
		beego.NSRouter("index", &controllers.QtIndexController{}, "GET:IndexPage"),
		//文章列表页
		beego.NSRouter("art", &controllers.QtIndexController{}, "GET:ArticleListPage"),
	)
	beego.AddNamespace(ns)
}
