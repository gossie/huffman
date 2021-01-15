package huffman

import "strings"

// BitSet is a bitset.
type BitSet struct {
	data []byte
}

// From creates a new bitset from the given slice.
func From(data []byte) BitSet {
	return BitSet{data: data}
}

func (b BitSet) String() string {
	stringBuilder := strings.Builder{}
	for _, bytes := range b.data {
		mask := byte(1)
		for i := 0; i < 7; i++ {
			if (bytes & mask) != 0 {
				stringBuilder.WriteString("1")
			} else {
				stringBuilder.WriteString("0")
			}
			mask = mask << 1
		}
	}
	return stringBuilder.String()
}

// IsSet returns true if the bit at the given index is set.
func (b *BitSet) IsSet(index uint) bool {
	dataIndex := index >> 3
	if dataIndex >= uint(len(b.data)) {
		return false
	}
	value := 1 << (index & 7)
	return (b.data[dataIndex] & byte(value)) > 0
}

// Set sets the bit at the given index.
func (b *BitSet) Set(index uint) {
	dataIndex := index >> 3
	for i := uint(len(b.data)); i <= dataIndex; i++ {
		b.data = append(b.data, 0)
	}
	value := 1 << (index & 7)
	b.data[dataIndex] |= byte(value)
}

// Bytes returns the a slice of byte.
func (b *BitSet) Bytes() []byte {
	return b.data
}
