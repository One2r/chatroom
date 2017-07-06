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

package models

import (
	"container/list"
)

var (
	//Subscribers 聊天室在线连接
	Subscribers = make(map[int]*list.List)

	//Roomconf 聊天室相关配置
	Roomconf = make(map[int]*RoomConf)
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
	EVENT_BIZ_EXCEPTION
)

type Event struct {
	Type      EventType // JOIN, LEAVE, MESSAGE
	ClientId  string
	Timestamp int // Unix timestamp (secs)
	Content   string
	Room      int
}

//RoomConf 聊天室配置
type RoomConf struct {
	MaxOnline       int          //最大在线人数
	Silence         bool         //(除管理员)全员禁言
	SpeakNotAllowed map[int]bool //被禁言人员
}

//User 用户信息
type User struct {
	ID       int    //用户ID
	Type     string //用户类型
	Username string //用户昵称
}
