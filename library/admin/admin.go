package admin

import (
	"chatroom/models"

	"github.com/astaxie/beego"
)

//GetStatis 获取聊天室相关统计信息
func GetStatis() map[string]interface{} {
	statis := make(map[string]interface{})
	statis["online"] = 0
	statis["roomNum"] = len(models.Subscribers)
	for _, room := range models.Subscribers {
		statis["online"] = statis["online"].(int) + room.Len()
	}
	beego.Info(statis)
	return statis
}
