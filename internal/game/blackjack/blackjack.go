package blackjack

import (
	"encoding/json"
	"log"

	"github.com/sam-vermeulen/go-poker/internal/types"
)

const (
	BlackjackHit   types.MessageType = "BLACKJACK_HIT"
	BlackjackStand types.MessageType = "BLACKJACK_STAND"
)

type BlackjackGame struct {
	Players map[string]*types.Player
	State   string
}

func NewBlackjackGame() *BlackjackGame {
	blackjack := &BlackjackGame{
		Players: make(map[string]*types.Player),
	}

	return blackjack
}

func (g *BlackjackGame) AddPlayer(player *types.Player) {
	log.Printf("Added player %s to blackjack", player.Name)
	g.Players[player.Name] = player
}

func (g *BlackjackGame) HandleAction(player *types.Player, action types.Message) error {
	if g.IsValidAction(player, action) {
		log.Printf("%s performed action %s", player.Name, action.Payload)
	}
	return nil
}

func (g *BlackjackGame) Start() error {

	return nil
}

func (g *BlackjackGame) GetState() interface{} {
	var state struct {
	}
	return state
}

func (g *BlackjackGame) IsValidAction(player *types.Player, action types.Message) bool {

	if _, exists := g.Players[player.Name]; !exists {
		log.Printf("player not exist")
		return false
	}

	var actionData struct {
		Type types.MessageType `json:"type"`
	}

	if err := json.Unmarshal(action.Payload, &actionData); err != nil {
		log.Printf("failed unmsarhal")
		return false
	}

	switch actionData.Type {
	case BlackjackHit:
		return true
	case BlackjackStand:
		return true
	default:
		return false
	}
}
