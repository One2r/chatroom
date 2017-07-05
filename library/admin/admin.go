package admin

import (
	"chatroom/models"
)

//GetStatis 获取聊天室相关统计信息
func GetStatis() map[string]interface{} {
	statis := make(map[string]interface{})
	statis["online"] = 0
	statis["MaxOnline"] = 0
	statis["roomNum"] = len(models.Subscribers)
	if statis["roomNum"].(int) > 0 {
		rooms := make(map[int]map[string]int)
		for k, room := range models.Subscribers {
			statis["online"] = statis["online"].(int) + room.Len()
			statis["MaxOnline"] = statis["MaxOnline"].(int) + models.Roomconf[k].MaxOnline
			roomInfo := make(map[string]int)
			roomInfo["online"] = room.Len()
			if models.Roomconf[k].Silence {
				roomInfo["Silence"] = 1
			} else {
				roomInfo["Silence"] = 0
			}
			roomInfo["MaxOnline"] = models.Roomconf[k].MaxOnline
			rooms[k] = roomInfo
		}
		statis["rooms"] = rooms
	}
	return statis
}
