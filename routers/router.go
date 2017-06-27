package routers

import (
	"github.com/astaxie/beego"

	admin "chatroom/app/admin/controllers"
	chat "chatroom/app/chat/controllers"
)

func init() {
	//默认路由为websocket连接地址
	beego.Router("/", &chat.WebSocketController{}, "get:Join")

	//chatroom后台管理接口路由
	adminNS := beego.NewNamespace("/admin",
		beego.NSRouter("/sensitive/update", &admin.AdminController{}, "get:UpdateSensitiveWords"),
		beego.NSRouter("/replace/update", &admin.AdminController{}, "get:UpdateReplaceWords"),
	)
	beego.AddNamespace(adminNS)

	//chatroom后台房管接口路由
	managerNS := beego.NewNamespace("/openapi",
		beego.NSRouter("/room/silence", &admin.ManagerController{}, "get:SetRoomSilence"),
	)
	beego.AddNamespace(managerNS)
}
