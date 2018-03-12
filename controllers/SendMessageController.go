package controllers

import (
	"ahaoouba/models"
)

func init() {
	go handleMessages()
	go talkroomMessages()
}
func handleMessages() {
	for {
		msg := <-models.Broadcast
		for k, ws := range models.Wss {
			if ws.UserName == msg.Message.Fsuname || ws.UserName == msg.Message.Jsuname {
				err := ws.WsConn.WriteJSON(msg)
				if err != nil {
					ws.WsConn.Close()
					models.Wss = append(models.Wss[:k], models.Wss[k+1:]...)
				}
			}

		}

	}

}
func talkroomMessages() {
	for {
		msg := <-models.RoomBroadcast
		for k, ws := range models.RoomWss {
			if ws.RoomCode == msg.Message.Code {
				err := ws.WsConn.WriteJSON(msg)
				if err != nil {
					ws.WsConn.Close()
					models.RoomWss = append(models.RoomWss[:k], models.RoomWss[k+1:]...)
				}
			}

		}

	}
}
