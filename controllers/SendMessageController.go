package controllers

import (
	"ahaoouba/models"
)

func init() {
	go handleMessages()
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
