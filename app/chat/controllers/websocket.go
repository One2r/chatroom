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

	"github.com/astaxie/beego"
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
	// Join chat room.
	Join(clientId, ws, room)
	defer func() {
		Leave(clientId, room)
		http.Error(this.Ctx.ResponseWriter, "连接已断开", 400)
		return
	}()

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		msg := string(p)

		if models.Roomconf[room].Silence { //全员禁言中
			publish <- newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "管理员开启了全员禁言", room)
			continue
		}

		if uSpeakNotAllowed, ok := models.Roomconf[room].SpeakNotAllowed[int(user.ID)]; ok && uSpeakNotAllowed { //个人被禁言
			publish <- newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "您被管理员禁言了", room)
			continue
		}

		if sensitive.Enable && sensitive.HasSensitiveWords(msg) { //敏感词信息屏蔽
			publish <- newEvent(models.EVENT_BIZ_EXCEPTION, clientId, "您的发言含有被屏蔽的关键词", room)
			continue
		}

		if replace.Enable { //内容替换
			msg = replace.Replace(msg)
		}

		publish <- newEvent(models.EVENT_MESSAGE, clientId, msg, room)
	}
}

// broadcastWebSocket broadcasts messages to WebSocket users.
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := models.Subscribers[event.Room].Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			switch event.Type {
			case models.EVENT_JOIN, models.EVENT_BIZ_EXCEPTION: //EVENT_JOIN EVENT_BIZ_EXCEPTION事件消息只发送给当前连接
				if event.ClientId == sub.Value.(Subscriber).ClientId {
					if ws.WriteMessage(websocket.TextMessage, data) != nil {
						// User disconnected.
						unsubscribe <- UnSubscriber{ClientId: sub.Value.(Subscriber).ClientId, Room: sub.Value.(Subscriber).Room}
					}
				}
				break
			default:
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
					// User disconnected.
					unsubscribe <- UnSubscriber{ClientId: sub.Value.(Subscriber).ClientId, Room: sub.Value.(Subscriber).Room}
				}
				break
			}
		}
	}
}
