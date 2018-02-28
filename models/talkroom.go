package models

import (
	"github.com/gorilla/websocket"
)

type TalkRoom struct {
	RoomCode string
	Ip       string
	WsConn   *websocket.Conn
}

var (
	RoomWss       = make([]*TalkRoom, 0)
	RoomBroadcast = make(chan *RoomTalkData)
)

type RoomTalkData struct {
	Message *RoomTalk
}
type RoomTalk struct {
	Fsname  string
	Content string
	Fstime  string
	Code    string
	Type    int64
}

const (
	HuanYing int64 = 1
	LiaoTian int64 = 2
)
