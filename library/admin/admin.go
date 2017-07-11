package admin

import (
	"chatroom/models"
	"strings"

	"strconv"

	"github.com/garyburd/redigo/redis"
)

//GetStatis 获取聊天室相关统计信息
func GetStatis() map[string]interface{} {
	redisConn := models.RedisConnPool.Get()
	chatRoom, _ := redis.Strings(redisConn.Do("PUBSUB", "CHANNELS", "chat_room_*"))
	statis := make(map[string]interface{})
	statis["online"] = 0
	statis["MaxOnline"] = 0
	statis["roomNum"] = len(chatRoom)
	if statis["roomNum"].(int) > 0 {
		rooms := make(map[int]map[string]int)
		for _, room := range chatRoom {
			roomArr := strings.Split(room, "_")
			roomConfNS := "RoomConfig:" + roomArr[2] + ":"

			Online, _ := redis.IntMap(redisConn.Do("PUBSUB", "NUMSUB", room))
			statis["online"] = statis["online"].(int) + Online[room]
			MaxOnline, _ := redis.Int(redisConn.Do("GET", roomConfNS+"MaxOnline"))
			statis["MaxOnline"] = statis["MaxOnline"].(int) + MaxOnline

			roomInfo := make(map[string]int)
			roomInfo["online"] = Online[room]
			Silence, _ := redis.Bool(redisConn.Do("GET", roomConfNS+"Silence"))
			if Silence {
				roomInfo["Silence"] = 1
			} else {
				roomInfo["Silence"] = 0
			}
			roomInfo["MaxOnline"] = MaxOnline
			roomk, _ := strconv.Atoi(roomArr[2])
			rooms[roomk] = roomInfo

		}
		statis["rooms"] = rooms
	}
	return statis
}
