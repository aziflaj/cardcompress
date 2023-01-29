package cardistry

import (
	"math/rand"
	"time"
)

type Deck []Card // Deck is a slice of Card

func NewDeck() *Deck {
	d := &Deck{}
	suits := []string{Spades, Hearts, Diamonds, Clubs}
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
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

// Compress the deck
// @return sign: false means the series starts with Black,
//
//	true means it starts with Red
//
// @return sequences: lengs of the sequences of the same color
func (d *Deck) Compress() (bool, []uint8) {
	sign := false
	firstCard := (*d)[0]
	if firstCard.Red() {
		sign = true
	}

	// count them
	var sequences []uint8
	prevIndex := 0
	for index, card := range *d {
		if index == 0 { // First count, nothing to compare with
			sequences = append(sequences, 1)
			continue
		}

		// Compare with previous card
		prevCard := (*d)[index-1]

		// If the color is the same, increment the count
		if card.Color() == prevCard.Color() {
			sequences[prevIndex] += 1
			continue
		}

		// add a new count
		sequences = append(sequences, 1)
		prevIndex += 1
	}

	return sign, sequences
}

func (d *Deck) String() string {
	var s string
	for index, card := range *d {
		s += card.String() + " "
		if (index+1)%6 == 0 {
			s += "\n"
		}
	}
	return s
}
