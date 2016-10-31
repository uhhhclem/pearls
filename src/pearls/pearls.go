package pearls

import (
	"log"
)

const maxNum = 100000

type Bitmap []byte

func NewBitmap() Bitmap {
	b := make([]byte, maxNum/8, maxNum/8)
	return Bitmap(b)
}

func (b Bitmap) SetBit(n int) {
	if n < 0 || n > maxNum {
		log.Fatalf("SetBit:  invalid argument: %d", n)
	}
	i, mask := getIndexAndMask(n)
	b[i] = b[i] | mask
}

func (b Bitmap) GetBit(n int) bool {
	if n < 0 || n > maxNum {
		log.Fatalf("GetBit:  invalid argument: %d", n)
	}
	i, mask := getIndexAndMask(n)
	return (b[i] & mask != 0)
}

func getIndexAndMask(n int) (int, byte) {
	i := n / 8
	bit := byte(n % 8)
	mask := byte(1 << bit)
	return i, mask	
}