package huffman

import (
	"testing"
)

func TestBitset(t *testing.T) {
	b := bitset{}
	b.set(0)
	b.set(2)
	if !b.isSet(0) {
		t.Fail()
	}
	if b.isSet(1) {
		t.Fail()
	}
	if !b.isSet(2) {
		t.Fail()
	}
}
