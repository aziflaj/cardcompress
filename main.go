package main

import (
	"aziflaj/cardcompress/cardistry"
	"fmt"
	"strconv"
)

func main() {
	// Deck magistry
	deck := cardistry.NewDeck()

	matrix := NewDeckMatrix(deck)

	fmt.Println(matrix)
}

type DeckMatrix struct {
	Sign  bool
	Frame []uint32
}

// Creates new DeckMatrix from a deck
// @param d: the deck to be compressed
// @return DeckMatrix: the compressed deck, with a Sign and a Frame
func NewDeckMatrix(d *cardistry.Deck) *DeckMatrix {
	matrix := &DeckMatrix{}
	sign, arr := d.Compress()
	matrix.Sign = sign

	robin := 0
	bigboi := uint32(0)
	for _, num := range arr {
		bigboi = bigboi | uint32(num)<<(robin*8)
		robin++

		if robin == 4 { // reset robin
			matrix.Frame = append(matrix.Frame, bigboi)
			robin = 0
			bigboi = 0
		}
	}

	return matrix
}

func (cd *DeckMatrix) String() string {
	var s string
	s += fmt.Sprintf("Sign: %v\n", cd.Sign)

	s += "Frame: ["
	for _, num := range cd.Frame {
		// s += strconv.FormatInt(int64(num), 16)
		s += strconv.Itoa(int(num))
		s += " "
	}
	s += fmt.Sprint("\b]\n")

	return s
}
