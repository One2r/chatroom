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
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/websocket"

	"chatroom/library/auth"
	"chatroom/library/filters/replace"
	"chatroom/library/filters/sensitive"
	"chatroom/models"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	beego.Controller
}

// Join method handles WebSocket requests for WebSocketController.
func (this *WebSocketController) Join() {
	tokenStr := this.GetString("token")
	if tokenStr == "" {
		http.Error(this.Ctx.ResponseWriter, "缺失连接token", 400)
		return
	}
	user, err := auth.CheckToken(tokenStr)
	if err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), 400)
		return
	}
	room, err := this.GetInt("room")
	if room <= 0 && err != nil {
		http.Error(this.Ctx.ResponseWriter, "房间号错误", 400)
		return
	}

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	clientId := NewClientId(room, ws.RemoteAddr().String())

	// send Join chat room msg.
	models.IsRoomConfigInit(room)
	if send(ws, newEvent(models.EVENT_JOIN, clientId, "", room)) != nil {
		http.Error(this.Ctx.ResponseWriter, "连接已断开", 400)
		return
	}

	psc := models.Subscribe(room)
	redisConn := models.RedisConnPool.Get()

	defer func() {
		redisConn.Close()
		models.UnSubscribe(psc)
		http.Error(this.Ctx.ResponseWriter, "连接已断开", 400)
		return
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			switch n := psc.Receive().(type) {
			case redis.Message:
				if send(ws, newEvent(models.EVENT_MESSAGE, clientId, string(n.Data), room)) != nil {
					return
				}
			case error:
				beego.Info(n)
				return
			}
		}
	}()

	// Message receive loop .
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		msg := string(p)

		//全员禁言中
		Silence, _ := redis.Bool(redisConn.Do("GET", "RoomConfig:"+strconv.Itoa(room)+":Silence"))
		if Silence {
			if send(ws, newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "管理员开启了全员禁言", room)) != nil {
				return
			}
			continue
		}

		//个人被禁言
		uSpeakNotAllowed, _ := redis.Bool(redisConn.Do("GET", "RoomConfig:"+strconv.Itoa(room)+":SpeakNotAllowed:"+strconv.Itoa(user.ID)))
		if uSpeakNotAllowed {
			send(ws, newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "您被管理员禁言了", room))
			continue
		}

		//敏感词信息屏蔽
		if sensitive.Enable && sensitive.HasSensitiveWords(msg) {
			if send(ws, newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "您的发言含有被屏蔽的关键词", room)) != nil {
				return
			}
			continue
		}

		//内容替换
		if replace.Enable {
			msg = replace.Replace(msg)
		}
		broadcastWebSocket(room, msg, redisConn)
	}
	wg.Wait()
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(room int, data string, redisConn redis.Conn) {
	redisConn.Send("PUBLISH", "chat_room_"+strconv.Itoa(room)+"_channel", data)
	redisConn.Flush()
}

//send 发送消息给websocket客户端
func send(ws *websocket.Conn, event models.Event) (errorVar error) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}
	errorVar = ws.WriteMessage(websocket.TextMessage, data)
	if errorVar != nil {
		ws.Close()
		beego.Error("WebSocket closed:", event.ClientId)
	}
	return errorVar
}
