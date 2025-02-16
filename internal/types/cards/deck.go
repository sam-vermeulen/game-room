package cards

import "math/rand/v2"

type Deck struct {
	cards   [52]uint
	topCard uint
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.initDeck()

	return deck
}

func (d *Deck) DrawCard() uint {
	card := d.cards[d.topCard]
	d.topCard += 1
	return card
}

func (d *Deck) ShuffleDeck() {
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *Deck) initDeck() {
	var n uint = 0
	var suit uint = 0x8000
	var i, j uint

	for i = 0; i < 4; i++ {
		for j = 0; j < 13; j++ {
			d.cards[n] = primes[j] | (j << 8) | suit | (1 << (16 + j))
			n++
		}
		suit >>= 1
	}
}

var primes []uint = []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
