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
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	"chatroom/models"
)

func newEvent(ep models.EventType, user, msg string, room int) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg, room}
}

func NewClientId(room int, RemoteAddr string) string {
	return "h_l_" + strconv.Itoa(room) + "_" + hex.EncodeToString([]byte(RemoteAddr))
}

func init() {
	go cleanEmptyRoom()
}

//清除空聊天室数据
func cleanEmptyRoom() {
	cleanTime := time.NewTicker(1 * 60 * time.Second)
	redisConn := models.RedisConnPool.Get()
	for {
		select {
		case <-cleanTime.C:
			chatRoom, _ := redis.Strings(redisConn.Do("KEYS", "RoomConfig*CreatedAt"))
			for _, room := range chatRoom {
				roomArr := strings.Split(room, ":")
				roomChannel := "chat_room_" + roomArr[1] + "_channel"
				Online, _ := redis.IntMap(redisConn.Do("PUBSUB", "NUMSUB", roomChannel))
				if Online[roomChannel] == 0 {
					EmptiedAt, _ := redis.Int(redisConn.Do("GET", "RoomConfig:"+roomArr[1]+":EmptiedAt"))
					if EmptiedAt > 0 && int(time.Now().Unix())-EmptiedAt > 60*30 {
						roomConfig, _ := redis.Strings(redisConn.Do("KEYS", "RoomConfig:"+roomArr[1]+":*"))
						for _, v := range roomConfig {
							redisConn.Do("DEL", v)
						}
					} else {
						redisConn.Do("SETNX", "RoomConfig:"+roomArr[1]+":EmptiedAt", int(time.Now().Unix()))
					}
				} else {
					redisConn.Do("DEL", "RoomConfig:"+roomArr[1]+":EmptiedAt")
				}
			}
		}
	}
}
