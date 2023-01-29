package cardistry

import (
	"fmt"
)

type Card struct {
	Number uint8
	Suit   string
}

func (c *Card) Color() string {
	if c.Suit == "♤" || c.Suit == "♧" {
		return "Black"
	}

	return "Red"
}

func (c *Card) Black() bool {
	return c.Suit == "♤" || c.Suit == "♧"
}

func (c *Card) Red() bool {
	return c.Suit == "♡" || c.Suit == "♢"
}

func (c *Card) String() string {
	if c.Number == 1 {
		return fmt.Sprintf("A%s", c.Suit)
	} else if c.Number == 11 {
		return fmt.Sprintf("J%s", c.Suit)
	} else if c.Number == 12 {
		return fmt.Sprintf("Q%s", c.Suit)
	} else if c.Number == 13 {
		return fmt.Sprintf("K%s", c.Suit)
	}

	return fmt.Sprintf("%d%s", c.Number, c.Suit)
}
