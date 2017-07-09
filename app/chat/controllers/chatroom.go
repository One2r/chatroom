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

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"chatroom/models"
)

func newEvent(ep models.EventType, user, msg string, room int) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg, room}
}

func Join(clientId string, ws *websocket.Conn, room int) {
	subscribe <- Subscriber{ClientId: clientId, Conn: ws, Room: room}
}

func Leave(clientId string, room int) {
	unsubscribe <- UnSubscriber{ClientId: clientId, Room: room}
}

func NewClientId(room int, RemoteAddr string) string {
	return "h_l_" + strconv.Itoa(room) + "_" + hex.EncodeToString([]byte(RemoteAddr))
}

type Subscriber struct {
	ClientId string
	Conn     *websocket.Conn // Only for WebSocket users; otherwise nil.
	Room     int
}

type UnSubscriber struct {
	ClientId string
	Room     int
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 50)
	// Channel for exit users.
	unsubscribe = make(chan UnSubscriber, 50)
	// Send events here to publish them.
	publish = make(chan models.Event, 100)
)

var rwmutex *sync.RWMutex

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if isRoomExist(models.Subscribers, sub.Room) {
				models.Subscribers[sub.Room].PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				publish <- newEvent(models.EVENT_JOIN, sub.ClientId, "", sub.Room)

				setRoomMaxOnline(sub.Room)
				beego.Info("New user:", sub.ClientId, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			broadcastWebSocket(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.ClientId, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:
			for sub := models.Subscribers[unsub.Room].Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).ClientId == unsub.ClientId {
					models.Subscribers[unsub.Room].Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub.ClientId)
					}
					break
				}
			}
		}
	}
}

func init() {
	rwmutex = new(sync.RWMutex)
	//go chatroom()
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
