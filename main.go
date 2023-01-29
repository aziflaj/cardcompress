package main

import (
	"aziflaj/cardcompress/cardistry"

	"fmt"
)

// Compress the deck
// @return sign: false means the series starts with Black,
//
//	true means it starts with Red
//
// @return compressArray: the array of the compressed deck
func compress(deck *cardistry.Deck) (bool, []uint8) {
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

	return sign, compressArray
}

func main() {
	// experiments with bit shifting
	// bitsy := 0b0010
	// fmt.Println(bitsy)
	// fmt.Println(bitsy << 1)
	// fmt.Println(bitsy << 2)

	sum := 0
	ca := 0

	// Deck magistry
	deck := cardistry.NewDeck()

	const trials = 100000
	// do the magic many times
	for i := 0; i < trials; i++ {
		deck.Shuffle()
		_, compressArray := compress(deck)

		caSum := len(compressArray)
		sum += caSum

		if i%100 == 0 {
			fmt.Println("CumAvg", ca)
			fmt.Println("Count", i)
			fmt.Println("Length of last run", len(compressArray))

			ca = ca + (caSum-ca)/(i+1)
			fmt.Println("CumAvg", ca)
		} else {
			ca = ca + (caSum-ca)/(i+1)
		}

	}

	fmt.Println("CumAvg", ca)
	fmt.Println("Sum", sum)
	fmt.Println("Avg", sum/trials)

}
