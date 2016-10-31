package pearls

const maxNum = 100000

type Bitmap []byte

func NewBitmap() Bitmap {
	b := make([]byte, maxNum/8, maxNum/8)
	return Bitmap(b)
}

func (b *Bitmap) SetBit(_ int) {

}

func (b *Bitmap) GetBit(_ int) bool {
	return false
}