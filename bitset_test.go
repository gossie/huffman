package huffman

import (
	"testing"
)

func TestBitset(t *testing.T) {
	b := BitSet{}
	b.Set(0)
	b.Set(2)
	if !b.IsSet(0) {
		t.Fail()
	}
	if b.IsSet(1) {
		t.Fail()
	}
	if !b.IsSet(2) {
		t.Fail()
	}
}

func TestBoundaries(t *testing.T) {
	b := BitSet{}
	b.Set(0)
	b.Set(7)
	if !b.IsSet(0) {
		t.Fail()
	}
	if b.IsSet(1) {
		t.Fail()
	}
	if b.IsSet(2) {
		t.Fail()
	}
	if b.IsSet(3) {
		t.Fail()
	}
	if b.IsSet(4) {
		t.Fail()
	}
	if b.IsSet(5) {
		t.Fail()
	}
	if b.IsSet(6) {
		t.Fail()
	}
	if !b.IsSet(7) {
		t.Fail()
	}
}

func TestBitsetBiggerThan64Entries(t *testing.T) {
	b := BitSet{}
	b.Set(67)
	b.Set(69)

	for i := uint(0); i < 67; i++ {
		if b.IsSet(i) {
			t.Fail()
		}
	}

	if !b.IsSet(67) {
		t.Fail()
	}
	if b.IsSet(68) {
		t.Fail()
	}
	if !b.IsSet(69) {
		t.Fail()
	}
}

func TestBitsetBiggerThan64EntriesWithCopy(t *testing.T) {
	old := BitSet{}
	old.Set(67)
	old.Set(69)

	b := From(old.Bytes())

	for i := uint(0); i < 67; i++ {
		if b.IsSet(i) {
			t.Fail()
		}
	}

	if !b.IsSet(67) {
		t.Fail()
	}
	if b.IsSet(68) {
		t.Fail()
	}
	if !b.IsSet(69) {
		t.Fail()
	}
}
