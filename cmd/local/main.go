package main

import (
	"fmt"

	"github.com/sam-vermeulen/go-poker/internal/types/cards"
)

func main() {
	deck := cards.NewDeck()
	for range 52 {
		card := deck.DrawCard()
		fmt.Println(cards.CardToString(card))
	}
}
