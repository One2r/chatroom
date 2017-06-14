package routers

import (
	"github.com/astaxie/beego"

	"chatroom/controllers"
)

func init() {
    beego.Router("/", &controllers.WebSocketController{}, "get:Join")
}