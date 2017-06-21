package routers

import (
	"github.com/astaxie/beego"

	"chatroom/controllers"
)

func init() {
	//默认路由为websocket连接地址
    beego.Router("/", &controllers.WebSocketController{}, "get:Join")

	//api路由
	beego.Router("/api/sensitive/update", &controllers.ApiController{}, "get:UpdateSensitiveWords")
}