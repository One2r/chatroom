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
	"time"
	"strconv"
	"encoding/hex"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"chatroom/models"
)

type Subscription struct {
	Archive []models.Event      // All the events from the archive.
	New     <-chan models.Event // New events coming in.
}

func newEvent(ep models.EventType, user, msg string,room int) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg,room}
}

func Join(user string, ws *websocket.Conn,room int) {
	subscribe <- Subscriber{Name: user, Conn: ws,Room: room}
}

func Leave(user string,room int) {
	unsubscribe <- UnSubscriber{Name:user,Room:room}
}

func NewClientId(room int,RemoteAddr string) string {
	return "h_l_" + strconv.Itoa(room) + "_" + hex.EncodeToString([]byte(RemoteAddr))
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
	Room int
}

type UnSubscriber struct {
	Name string
	Room int
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// Channel for exit users.
	unsubscribe = make(chan UnSubscriber, 10)
	// Send events here to publish them.
	publish = make(chan models.Event, 10)

	subscribers = make(map[int]*list.List)
)

// This function handles all incoming chan messages.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name,sub.Room) {
				subscribers[sub.Room].PushBack(sub) // Add user to the end of list.
				// Publish a JOIN event.
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "",sub.Room)
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			broadcastWebSocket(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers[unsub.Room].Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub.Name {
					subscribers[unsub.Room].Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub.Name)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub.Name, "",unsub.Room) // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers map[int]*list.List, user string,room int) bool {
	if subscribers[room] != nil {
		for sub := subscribers[room].Front(); sub != nil; sub = sub.Next() {
			if sub.Value.(Subscriber).Name == user {
				return true
			}
		}
	}else{
		subscribers[room] = list.New()
	}
	return false
}
