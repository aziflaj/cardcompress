package main

import (
	"aziflaj/cardcompress/cardistry"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// Deck magistry
	deck := cardistry.NewDeck()
	deck.Shuffle()
	fmt.Println(deck)

	sign, tally := deck.Compress()
	fmt.Println(tally)
	// err := writeToFile("tally.bin", tally)
	// if err != nil {
	// 	panic(err)
	// }
	matrix := cardistry.NewColorSeq(sign, tally)
	fmt.Println(matrix)

	// err = matrix.Dump("matrix.bin")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println(matrix.Decompress())
}

func writeToFile(filename string, data any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	err = binary.Write(file, binary.LittleEndian, data)
	if err != nil {
		return err
	}

	file.Write(buf.Bytes())

	return nil
}
