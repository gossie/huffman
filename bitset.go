package huffman

import "math"

type bitset struct {
	data         []uint64
	lastSetIndex uint
}

func from(data []uint64) bitset {
	return bitset{data: data}
}

func (b *bitset) isSet(index uint) bool {
	dataIndex := index / 64
	if dataIndex >= uint(len(b.data)) {
		return false
	}
	value := math.Pow(2, float64(index%64))
	return (b.data[dataIndex] & uint64(value)) > 0
}

func (b *bitset) set(index uint) {
	dataIndex := index / 64
	for i := uint(len(b.data)); i <= dataIndex; i++ {
		b.data = append(b.data, 0)
	}
	value := math.Pow(2, float64(index%64))
	b.data[dataIndex] |= uint64(value)
}

func (b *bitset) bytes() []uint64 {
	return b.data
}
