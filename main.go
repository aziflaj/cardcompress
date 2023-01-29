package main

import (
	"aziflaj/cardcompress/cardistry"
	"fmt"
)

func main() {
	// Deck magistry
	deck := cardistry.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)

	sign, tally := deck.Compress()
	matrix := cardistry.NewDeckMatrix(sign, tally)
	fmt.Println(matrix)

	fmt.Println(matrix.Decompress())
}
