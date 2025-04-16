package blackjack

import (
	"math/rand/v2"

	"github.com/sam-vermeulen/go-poker/internal/types"
)

const (
	BlackjackHit   types.MessageType = "BLACKJACK_HIT"
	BlackjackStand types.MessageType = "BLACKJACK_STAND"
)

type GameState string

const (
	StateBetting  GameState = "betting"
	StatePlaying  GameState = "playing"
	StateSettling GameState = "settling"
	StateFinished GameState = "finished"
)

type Turn struct {
	Player    *types.Player `json:"player"`
	HandIndex int           `json:"handIndex"`
}

type PlayerState struct {
	Hands  []Hand `json:"hands"`
	Chips  int    `json:"chips"`
	Status string `json:"status"`
}

func (p *PlayerState) CanSplit(handIndex int) bool {
	if handIndex >= len(p.Hands) {
		return false
	}
	hand := p.Hands[handIndex]
	return len(hand.Cards) == 2 && p.Chips >= hand.Bet && !hand.Doubled
}

func (p *PlayerState) CanDouble(handIndex int) bool {
	if handIndex >= len(p.Hands) {
		return false
	}

	hand := p.Hands[handIndex]
	return len(hand.Cards) == 2 && p.Chips >= hand.Bet && !hand.Doubled
}

func newDeck() []Card {
	const numDecks = 6

	deck := make([]Card, 0, 52*numDecks)

	for d := 0; d < numDecks; d++ {
		for _, suit := range []Suit{Hearts, Diamonds, Clubs, Spades} {
			for _, value := range []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King} {
				deck = append(deck, Card{
					Suit:  suit,
					Value: value,
				})
			}
		}
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
