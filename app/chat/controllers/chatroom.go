// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

import (
	"container/list"
	"encoding/hex"
	"strconv"
	"sync"
	"time"

	"chatroom/models"
)

func newEvent(ep models.EventType, user, msg string, room int) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg, room}
}

func NewClientId(room int, RemoteAddr string) string {
	return "h_l_" + strconv.Itoa(room) + "_" + hex.EncodeToString([]byte(RemoteAddr))
}

var rwmutex *sync.RWMutex

func init() {
	rwmutex = new(sync.RWMutex)
	go cleanEmptyRoom()
}

//检查房间是否存在
func isRoomExist(subscribers map[int]*list.List, room int) bool {
	if subscribers[room] == nil {
		subscribers[room] = list.New()
		models.Roomconf[room] = &models.RoomConf{MaxOnline: 0, Silence: false, SpeakNotAllowed: make(map[int]bool)}
	}
	return true
}

//清除空聊天室数据
func cleanEmptyRoom() {
	cleanTime := time.NewTicker(30 * 60 * time.Second)
	for {
		select {
		case <-cleanTime.C:
			for k, v := range models.Subscribers {
				if v.Len() == 0 {
					delete(models.Subscribers, k)
					delete(models.Roomconf, k)
				}
			}
		}
	}
}

//更新聊天室最大在线人数
func setRoomMaxOnline(room int) {
	rwmutex.Lock()
	roomNum := models.Subscribers[room].Len()
	if roomNum > models.Roomconf[room].MaxOnline {
		models.Roomconf[room].MaxOnline = roomNum
	}
	rwmutex.Unlock()
}
