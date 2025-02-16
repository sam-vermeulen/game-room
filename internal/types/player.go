package types

import (
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name string
	Conn *websocket.Conn
}

type PlayerConnection interface {
	RemovePlayer(*Player)
	HandleMessage(*Player, Message)
}

func (p *Player) Send(msg interface{}) error {
	return p.Conn.WriteJSON(msg)
}

func (p *Player) HandleConnection(r PlayerConnection) {
	defer func() {
		r.RemovePlayer(p)
		p.Conn.Close()
	}()

	for {
		var msg Message
		err := p.Conn.ReadJSON(&msg)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		r.HandleMessage(p, msg)
	}
}
