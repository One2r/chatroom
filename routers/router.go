package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	admin "chatroom/app/admin/controllers"
	chat "chatroom/app/chat/controllers"

	"chatroom/library/auth"
)

var FilterAdminLogin = func(ctx *context.Context) {
	token := ctx.Input.Cookie("token")
	admin, err := auth.CheckAdminToken(token)
	if err != nil {
		if ctx.Request.RequestURI != "/admin/signin.html" && ctx.Request.RequestURI != "/admin/signout.html" {
			ctx.Redirect(302, "/admin/signin.html")
		}
	} else {
		ctx.Input.SetData("isLogin", admin)
		if ctx.Request.RequestURI == "/admin/signin.html" {
			ctx.Redirect(302, "/admin/dashboard.html")
		}
	}
}

func init() {
	//默认路由为websocket连接地址
	beego.Router("/", &chat.WebSocketController{}, "get:Join")

	//chatroom后台管理接口路由
	adminNS := beego.NewNamespace("/admin",
		beego.NSRouter("/signin", &admin.AdminController{}, "get,post:Signin"),
		beego.NSRouter("/signout", &admin.AdminController{}, "get,post:Signout"),

		beego.NSBefore(FilterAdminLogin),
		beego.NSRouter("/dashboard", &admin.AdminController{}, "get:Dashboard"),
		beego.NSRouter("/service", &admin.AdminController{}, "get:Service"),
		beego.NSRouter("/sensitive/update", &admin.AdminController{}, "get:UpdateSensitiveWords"),
		beego.NSRouter("/replace/update", &admin.AdminController{}, "get:UpdateReplaceWords"),
	)
	beego.AddNamespace(adminNS)

	//chatroom后台房管接口路由
	managerNS := beego.NewNamespace("/openapi",
		beego.NSRouter("/room/silence", &admin.ManagerController{}, "get:SetRoomSilence"),
		beego.NSRouter("/room/speaknotallowed", &admin.ManagerController{}, "get:SpeakNotAllowed"),
	)
	beego.AddNamespace(managerNS)
}
