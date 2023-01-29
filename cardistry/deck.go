package cardistry

import (
	"math/rand"
	"time"
)

type Deck []Card // Deck is a slice of Card

func NewDeck() *Deck {
	d := &Deck{}
	suits := []string{"♤", "♡", "♢", "♧"}
	for _, suit := range suits {
		for i := 1; i <= 13; i++ {
			*d = append(*d, Card{Number: uint8(i), Suit: suit})
		}
	}

	return d
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	// Fisher Yates shuffle
	for i := len(*d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) String() string {
	var s string
	for _, card := range *d {
		s += card.String() + " "
	}
	return s
}
