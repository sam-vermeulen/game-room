package blackjack

type Card struct {
	Suit  Suit  `json:"suit"`
	Value Value `json:"value"`
}

type Suit string

const (
	Hearts   Suit = "hearts"
	Diamonds Suit = "diamonds"
	Clubs    Suit = "clubs"
	Spades   Suit = "spades"
)

type Value string

const (
	Ace   Value = "A"
	King  Value = "K"
	Queen Value = "Q"
	Jack  Value = "J"
	Ten   Value = "10"
	Nine  Value = "9"
	Eight Value = "8"
	Seven Value = "7"
	Six   Value = "6"
	Five  Value = "5"
	Four  Value = "4"
	Three Value = "3"
	Two   Value = "2"
)

type Hand struct {
	Cards    []Card `json:"cards"`
	Bet      int    `json:"bet"`
	Doubled  bool   `json:"doubled"`
	IsSplit  bool   `json:"isSplit"`
	Standing bool   `json:"standing"`
}

func (h *Hand) Value() int {
	value := 0
	aces := 0

	for _, card := range h.Cards {
		switch card.Value {
		case Ace:
			aces++
		case King, Queen, Jack, Ten:
			value += 10
		default:
			value += int(card.Value[0] - '0')
		}
	}

	for i := 0; i < aces; i++ {
		if value+11 <= 21 {
			value += 11
		} else {
			value += 1
		}
	}

	return value
}

func (h *Hand) IsBusted() bool {
	return h.Value() > 21
}

func (h *Hand) IsBlackjack() bool {
	return h.Value() == 21
}
