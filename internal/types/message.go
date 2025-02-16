package types

import "encoding/json"

type MessageType string

const (
	MessageChat       MessageType = "CHAT"
	MessageGameStart  MessageType = "GAME_START"
	MessageGameEnd    MessageType = "GAME_END"
	MessageGameCreate MessageType = "GAME_CREATE"
)

type Message struct {
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
