package cardistry

import (
	"fmt"
)

const (
	Spades   string = "♠️"
	Hearts          = "♥️"
	Diamonds        = "♦️"
	Clubs           = "♣️"
)

type Card struct {
	Number int32
	Suit   string
}

func (c *Card) Color() string {
	if c.Black() {
		return "Black"
	}

	return "Red"
}

func (c *Card) Black() bool {
	return c.Suit == Spades || c.Suit == Clubs
}

func (c *Card) Red() bool {
	return c.Suit == Hearts || c.Suit == Diamonds
}

func (c *Card) String() string {
	if c.Number == 1 {
		return fmt.Sprintf(" A%s ", c.Suit)
	} else if c.Number == 11 {
		return fmt.Sprintf(" J%s ", c.Suit)
	} else if c.Number == 12 {
		return fmt.Sprintf(" Q%s ", c.Suit)
	} else if c.Number == 13 {
		return fmt.Sprintf(" K%s ", c.Suit)
	}

	return fmt.Sprintf("%2d%s ", c.Number, c.Suit)
}
