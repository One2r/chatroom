package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	admin "chatroom/app/admin/controllers"
	chat "chatroom/app/chat/controllers"

	"chatroom/library/auth"
	"chatroom/models"
)

//CheckAdminLogin 检查后台管理员是否登录
var CheckAdminLogin = func(ctx *context.Context) {
	token := ctx.Input.Cookie("token")
	if token == "" {
		if ctx.Request.RequestURI == "/admin/signin.html" {
			return
		}
		ctx.Redirect(301, "/admin/signin.html")
		return
	}

	admin, err := auth.CheckToken(token)
	if err != nil || admin.Type != models.USER_TYPE_ADMIN {
		ctx.Redirect(301, "/admin/signin.html")
		return
	}

	ctx.Input.SetData("isLogin", admin)
	if ctx.Request.RequestURI == "/admin/signin.html" {
		ctx.Redirect(301, "/admin/dashboard.html")
		return
	}
}

//CheckOpneAPIAuth 检查后台openapi权限
var CheckOpneAPIAuth = func(ctx *context.Context) {
	var token string
	ctx.Input.Bind(&token, "token")
	if token == "" {
		ctx.Output.JSON(admin.BizException("缺失token", 600), false, false)
		return
	}
	manager, err := auth.CheckToken(token)
	if err != nil || manager.Type != models.USER_TYPE_MANAGER {
		ctx.Output.JSON(admin.BizException("无效token", 600), false, false)
		return
	}
}

func init() {
	//默认路由为websocket连接地址
	beego.Router("/", &chat.WebSocketController{}, "get:Join")

	//chatroom后台管理接口路由
	adminNS := beego.NewNamespace("/admin",
		beego.NSRouter("/signin", &admin.AdminController{}, "get,post:Signin"),
		beego.NSRouter("/signout", &admin.AdminController{}, "get,post:Signout"),

		beego.NSBefore(CheckAdminLogin),
		beego.NSRouter("/dashboard", &admin.AdminController{}, "get:Dashboard"),
		beego.NSRouter("/service", &admin.AdminController{}, "get:Service"),
		beego.NSRouter("/sensitive/update", &admin.AdminController{}, "get:UpdateSensitiveWords"),
		beego.NSRouter("/replace/update", &admin.AdminController{}, "get:UpdateReplaceWords"),
	)
	beego.AddNamespace(adminNS)

	//chatroom后台房管接口路由
	managerNS := beego.NewNamespace("/openapi",
		beego.NSBefore(CheckOpneAPIAuth),
		beego.NSRouter("/room/silence", &admin.ManagerController{}, "get:SetRoomSilence"),
		beego.NSRouter("/room/speaknotallowed", &admin.ManagerController{}, "get:SpeakNotAllowed"),
	)
	beego.AddNamespace(managerNS)
}
