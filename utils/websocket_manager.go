package utils

import (
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"log"
)

type WebsocketManage struct {

}

func New() *WebsocketManage {

	return &WebsocketManage{}
}

func (c *WebsocketManage) Handler() *neffos.Server{
	//websocket
	//接收消息回调
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	})

	//连接回调
	ws.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
	}


	return ws
}

