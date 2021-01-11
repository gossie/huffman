package huffman

// BitSet is a bitset.
type BitSet struct {
	data []uint64
}

// From creates a new bitset from the given slice.
func From(data []uint64) BitSet {
	return BitSet{data: data}
}

// IsSet returns true if the bit at the given index is set.
func (b *BitSet) IsSet(index uint) bool {
	dataIndex := index >> 6
	if dataIndex >= uint(len(b.data)) {
		return false
	}
	value := 1 << (index & 63)
	return (b.data[dataIndex] & uint64(value)) > 0
}

// Set sets the bit at the given index.
func (b *BitSet) Set(index uint) {
	dataIndex := index >> 6
	for i := uint(len(b.data)); i <= dataIndex; i++ {
		b.data = append(b.data, 0)
	}
	value := 1 << (index & 63)
	b.data[dataIndex] |= uint64(value)
}

// Bytes returns the a slice of uint64.
func (b *BitSet) Bytes() []uint64 {
	return b.data
}
