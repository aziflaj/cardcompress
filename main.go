package main

import (
	"aziflaj/cardcompress/cardistry"

	"fmt"
)

func main() {
	bitsy := 0b0010
	fmt.Println(bitsy)
	fmt.Println(bitsy << 1)
	fmt.Println(bitsy << 2)

	deck := cardistry.NewDeck()
	deck.Shuffle()

	fmt.Println(deck)

	// false means the series starts with Black,
	// true means it starts with Red
	sign := false
	firstCard := (*deck)[0]
	if firstCard.Suit == "♡" || firstCard.Suit == "♢" {
		sign = true
	}

	// count them
	var compressArray []uint8
	prevIndex := 0
	for index, card := range *deck {
		if index == 0 {
			compressArray = append(compressArray, 1)
		} else {
			prevCard := (*deck)[prevIndex]

			if card.Color() != prevCard.Color() {
				compressArray = append(compressArray, 1)
				prevIndex += 1
			} else {
				compressArray[prevIndex] += 1
			}
		}
	}

	if sign {
		fmt.Println("starts with Red")
	} else {
		fmt.Println("starts with Black")
	}

	fmt.Println(len(compressArray))
	fmt.Println(compressArray)
}
