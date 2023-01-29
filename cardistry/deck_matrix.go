package cardistry

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

type DeckMatrix struct {
	Sign  bool
	Frame []int32
}

// Creates new DeckMatrix from a deck
// @param d: the deck to be compressed
// @return DeckMatrix: the compressed deck, with a Sign and a Frame
// func NewDeckMatrix(d *Deck) *DeckMatrix {
func NewDeckMatrix(sign bool, arr []int32) *DeckMatrix {
	// sign, arr := d.Compress()
	var frame []int32 = make([]int32, len(arr)/4+1)

	robin := 0
	frameIdx := 0
	bigboi := int32(0)
	for _, num := range arr {
		bigboi = bigboi | int32(num)<<(robin*8)

		robin++
		if robin == 4 { // reset robin and bigboi
			frame[frameIdx] = bigboi
			robin = 0
			bigboi = 0
			frameIdx++
		}

		frame[frameIdx] = bigboi
	}

	return &DeckMatrix{Sign: sign, Frame: frame}
}

func (dm *DeckMatrix) Decompress() string {
	var s string
	if dm.Sign {
		s += "R "
	} else {
		s += "B "
	}

	for _, num := range dm.Frame {
		// num needs to be and'd with 0xFF to get the last 8 bits
		left := num & 0xFF
		middle_l := (num >> 8) & 0xFF
		middle_r := (num >> 16) & 0xFF
		right := (num >> 24) & 0xFF

		s += fmt.Sprintf("%d %d %d %d ", left, middle_l, middle_r, right)
	}

	return s
}

func (m *DeckMatrix) String() string {
	var sign string
	if m.Sign {
		sign = "R"
	} else {
		sign = "B"
	}

	var sFrame string
	for _, num := range m.Frame {
		sFrame += "0x" + strconv.FormatInt(int64(num), 16) + " "
		// sFrame += strconv.Itoa(int(num)) + " "
	}

	return fmt.Sprintf("DeckMatrix: {Sign: %v; Frame: [%s\b]}\n", sign, sFrame)
}

func (m *DeckMatrix) Dump(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	// err = binary.Write(buf, binary.LittleEndian, m.Sign)
	err = binary.Write(buf, binary.LittleEndian, m.Frame)
	if err != nil {
		return err
	}

	f.Write(buf.Bytes())
	return nil
}
