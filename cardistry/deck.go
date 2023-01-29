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

// Compress the deck
// @return sign: false means the series starts with Black,
//
//	true means it starts with Red
//
// @return compressArray: the array of the compressed deck
func (d *Deck) Compress() (bool, []uint8) {
	sign := false
	firstCard := (*d)[0]
	if firstCard.Suit == "♡" || firstCard.Suit == "♢" {
		sign = true
	}

	// count them
	var compressArray []uint8
	prevIndex := 0
	for index, card := range *d {
		if index == 0 {
			compressArray = append(compressArray, 1)
			continue
		}

		prevCard := (*d)[prevIndex]

		if card.Color() == prevCard.Color() {
			compressArray[prevIndex] += 1
			continue
		}

		compressArray = append(compressArray, 1)
		prevIndex += 1
	}

	return sign, compressArray
}

func (d *Deck) String() string {
	var s string
	for _, card := range *d {
		s += card.String() + " "
	}
	return s
}
