package blackjack

import (
	"encoding/json"
	"log"
	"time"

	"github.com/sam-vermeulen/go-poker/internal/types"
)

type BlackjackGame struct {
	Players     map[string]*PlayerState `json:"players"`
	Dealer      Hand                    `json:"dealer"`
	Deck        []Card                  `json:"-"`
	State       GameState               `json:"state"`
	Round       int                     `json:"round"`
	CurrentTurn *Turn                   `json:"currentTurn"`
	LastAction  time.Time               `json:"lastAction"`
}

func NewBlackjackGame() *BlackjackGame {
	blackjack := &BlackjackGame{
		Players: make(map[string]*PlayerState),
		Deck:    newDeck(),
		State:   StateBetting,
		Round:   1,
	}

	return blackjack
}

func (g *BlackjackGame) AddPlayer(player *types.Player) {
	log.Printf("Added player %s to blackjack", player.Name)
	g.Players[player.Name] = &PlayerState{
		Hands:  make([]Hand, 1),
		Chips:  1000,
		Status: "betting",
	}

	u, err := json.Marshal(g)
	if err != nil {
		log.Printf("marshall error: %v", err)
		return
	}
	log.Println(string(u))
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

func (g *BlackjackGame) drawCard() Card {
	if len(g.Deck) < 20 {
		g.Deck = newDeck()
	}

	card := g.Deck[len(g.Deck)-1]
	g.Deck = g.Deck[:len(g.Deck)-1]

	return card
}

func (g *BlackjackGame) GetPlayerState(player *types.Player) map[string]interface{} {
	p := g.Players[player.Name]
	return map[string]interface{}{
		"hands":       p.Hands,
		"chips":       p.Chips,
		"dealer":      g.Dealer,
		"state":       g.State,
		"currentTurn": g.CurrentTurn,
		"canSplit":    p.CanSplit(0),
		"canDouble":   p.CanDouble(0),
	}
}
