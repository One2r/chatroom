package manager

import (
	"chatroom/models"
	"strconv"
)

// SetRoomSilence 设置某个房间的全员禁言状态
func SetRoomSilence(room int, status string) bool {
	redisConn := models.RedisConnPool.Get()
	_, err := redisConn.Do("SET", "RoomConfig:"+strconv.Itoa(room)+":Silence", status)
	redisConn.Close()
	if err != nil {
		return false
	}
	return true
}

//SpeakNotAllowed 禁言某个房间的某个人
func SpeakNotAllowed(room int, uid int, status string) bool {
	if roomconf, ok := models.Roomconf[room]; ok {
		if status == "true" {
			roomconf.SpeakNotAllowed[uid] = true
		} else {
			roomconf.SpeakNotAllowed[uid] = false
		}
		return true
	}
	return false
}
