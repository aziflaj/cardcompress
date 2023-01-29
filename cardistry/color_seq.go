package cardistry

import (
	"fmt"
	"strconv"
)

type ColorSeq struct {
	Sign  bool
	Frame []uint32
}

// Creates new ColorSeq from a deck
// @param sign: the sign of the deck
// @param arr: array of sequence lengths
// @return ColorSeq: the compressed sequence
func NewColorSeq(sign bool, arr []uint8) *ColorSeq {
	frame := make([]uint32, len(arr)/6+1)

	robin := 0
	frameIdx := 0
	bigboi := uint32(0)
	for _, num := range arr {
		bigboi = bigboi | uint32(num)<<(robin*5)

		robin++
		if robin == 6 { // reset robin and bigboi
			frame[frameIdx] = bigboi
			robin = 0
			bigboi = 0
			frameIdx++
		}

		frame[frameIdx] = bigboi
	}

	return &ColorSeq{Sign: sign, Frame: frame}
}

// Convert from a ColorSeq to a sign and a tally
func (cs *ColorSeq) Decompress() string {
	var s string
	if cs.Sign {
		s += "R "
	} else {
		s += "B "
	}

	for _, num := range cs.Frame {
		// num needs to be and'd with 0x1F to get the last 5 bits
		for i := 0; i < 6; i++ {
			s += strconv.Itoa(int(num&0x1F)) + " "
			num = num >> 5
		}
	}

	return s
}

func (cs *ColorSeq) String() string {
	var sign string
	if cs.Sign {
		sign = "R"
	} else {
		sign = "B"
	}

	var sFrame string
	for _, num := range cs.Frame {
		sFrame += "0x" + strconv.FormatInt(int64(num), 16) + " "
		// sFrame += strconv.Itoa(int(num)) + " "
	}

	return fmt.Sprintf("ColorSeq: {Sign: %v; Frame: [%s\b]}\n", sign, sFrame)
}
