package game

import (
	"encoding/json"

	"github.com/sam-vermeulen/go-poker/internal/types"
)

type GamePlayer struct {
	Name string
}

type GameAction struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type GameState struct {
}

type GameType interface {
	HandleAction(player *types.Player, action types.Message) error
	Start() error
	GetState() interface{}
	IsValidAction(player *types.Player, action types.Message) bool
}
