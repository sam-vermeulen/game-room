package room

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/sam-vermeulen/go-poker/internal/game"
	"github.com/sam-vermeulen/go-poker/internal/game/blackjack"
	"github.com/sam-vermeulen/go-poker/internal/types"
)

type Room struct {
	mu           sync.RWMutex
	Code         string
	Players      map[string]*types.Player
	Game         game.GameType
	OnEmpty      func(code string)
	createdAt    time.Time
	cleanupTimer *time.Timer
	texts        []*types.Message
}

type TextMessage struct {
	Text      string    `json:"text"`
	Sender    string    `json:"sender,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func NewRoom(code string) *Room {
	room := &Room{
		Code:    code,
		Players: make(map[string]*types.Player),
	}

	return room
}

func (r *Room) IsFull() bool {
	return false
}

func (r *Room) resetCleanupTime(dur time.Duration) {
	if r.cleanupTimer != nil {
		r.cleanupTimer.Stop()
	}

	r.cleanupTimer = time.AfterFunc(dur, func() {
		r.mu.Lock()
		defer r.mu.Unlock()

		if len(r.Players) == 0 && r.OnEmpty != nil {
			r.OnEmpty(r.Code)
		}
	})
}

func (r *Room) AddPlayer(p *types.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Players[p.Name] = p
	log.Printf("Player %s has joined room %s", p.Name, r.Code)

	if r.cleanupTimer != nil {
		r.cleanupTimer.Stop()
	}

	for _, text := range r.texts {
		p.Send(text)
	}
}

func (r *Room) RemovePlayer(p *types.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Players, p.Name)
	log.Printf("Player %s has left room %s", p.Name, r.Code)

	if len(r.Players) == 0 {
		r.resetCleanupTime(20 * time.Second)
	}
}

func (r *Room) broadcastMessage(msg interface{}) {
	for _, player := range r.Players {
		go player.Send(msg)
	}
}

func (r *Room) selectGame(p *types.Player, selectedGame string) {
	log.Printf("Selecting game %s", selectedGame)
	switch selectedGame {
	case "blackjack":
		blackjack := blackjack.NewBlackjackGame()
		for _, player := range r.Players {
			blackjack.AddPlayer(player)
		}
		r.Game = blackjack
		return
	default:
		p.Send("Cannot create a new game, not a valid game")
	}
}

func (r *Room) HandleMessage(p *types.Player, msg types.Message) {
	switch msg.Type {
	case types.MessageChat:
		var text TextMessage

		text.Timestamp = time.Now()
		text.Sender = p.Name

		if err := json.Unmarshal(msg.Payload, &text); err != nil {
			go p.Send("Incorrect text format")
			return
		}

		newPayload, err := json.Marshal(text)
		if err != nil {
			go p.Send("Error creating message")
			return
		}

		msg.Payload = newPayload

		r.texts = append(r.texts, &msg)

		r.broadcastMessage(msg)
		log.Printf("Chat message to %s from %s: %s", r.Code, p.Name, text.Text)
		return
	case types.MessageGameCreate:
		if r.Game != nil {
			go p.Send("Cannot create a new game, game already selected")
			return
		}

		var gameSelect struct {
			Game string `json:"type"`
		}

		if err := json.Unmarshal(msg.Payload, &gameSelect); err != nil {
			go p.Send("Cannot create a new game, invalid payload")
			return
		}

		r.selectGame(p, gameSelect.Game)
		return
	default:
		if r.Game != nil {
			r.Game.HandleAction(p, msg)
		} else {
			go p.Send("No active game")
		}
		return
	}
}
