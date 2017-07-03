package manager

import (
	"chatroom/models"
)

// SetRoomSilence 设置某个房间的全员禁言状态
func SetRoomSilence(room int, status string) bool {
	if roomconf, ok := models.Roomconf[room]; ok {
		if status == "true" {
			roomconf.Silence = true
		} else {
			roomconf.Silence = false
		}
		return true
	} else {
		return false
	}
}
