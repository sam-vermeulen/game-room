package cards

const (
	// CARD SUIT
	CLUB    = 0x7fff
	DIAMOND = 0x3fff
	HEART   = 0x1fff
	SPADE   = 0x0fff

	// CARD VALUE
	TWO   = -1
	THREE = 0
	FOUR  = 1
	FIVE  = 2
	SIX   = 3
	SEVEN = 4
	EIGHT = 5
	NINE  = 6
	TEN   = 7
	JACK  = 8
	QUEEN = 9
	KING  = 10
	ACE   = 11
)

func GetRank(card uint) string {
	rank := (card >> 8) & 0x0F
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	return ranks[rank]
}

func GetSuit(card uint) string {
	suit := (card >> 12) & 0x0F
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}

	suitIndex := 0
	switch suit {
	case 0x8:
		suitIndex = 0 // clubs
	case 0x4:
		suitIndex = 1 // diamonds
	case 0x2:
		suitIndex = 2 // hearts
	case 0x1:
		suitIndex = 3 // spades
	}

	return suits[suitIndex]
}

func CardToString(card uint) string {
	return GetRank(card) + " of " + GetSuit(card)
}
