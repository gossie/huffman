package huffman

import (
	"testing"
)

func TestFromInputWithOnlyOneLetter(t *testing.T) {
	root := fromInput("a")
	if root.one != nil || root.zero != nil {
		t.Fatal("The node should not have children")
	}
}

func TestFromInputWithOneDistinctLetter(t *testing.T) {
	root := fromInput("aaa")
	if root.letter != "a" || root.frequency != 1.0 || root.one != nil || root.zero != nil {
		t.Fatalf(`node.letter = %v, node.frequency = %v`, root.letter, root.frequency)
	}
}

func TestFromInputWithTwoDistinctLetters(t *testing.T) {
	root := fromInput("aaba")
	if root.letter != "" || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`root.letter = %v, root.frequency = %v`, root.letter, root.frequency)
	}

	zero := root.zero
	if zero.letter != "b" || zero.frequency != 0.25 || zero.one != nil || zero.zero != nil {
		t.Fatalf(`zero.letter = %v, zero.frequency = %v`, zero.letter, zero.frequency)
	}

	one := root.one
	if one.letter != "a" || one.frequency != 0.75 || one.one != nil || one.zero != nil {
		t.Fatalf(`one.letter = %v, one.frequency = %v`, one.letter, one.frequency)
	}
}

func TestFromInputWithThreeDistinctLetters(t *testing.T) {
	root := fromInput("aabaabca")
	if root.letter != "" || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`root.letter = %v, root.frequency = %v`, root.letter, root.frequency)
	}

	zero := root.zero
	if zero.letter != "" || zero.frequency != 0.375 || zero.one == nil || zero.zero == nil {
		t.Fatalf(`zero.letter = %v, zero.frequency = %v`, zero.letter, zero.frequency)
	}

	zeroZero := zero.zero
	if zeroZero.letter != "c" || zeroZero.frequency != 0.125 || zeroZero.one != nil || zeroZero.zero != nil {
		t.Fatalf(`zeroZero.letter = %v, zeroZero.frequency = %v`, zeroZero.letter, zeroZero.frequency)
	}

	zeroOne := zero.one
	if zeroOne.letter != "b" || zeroOne.frequency != 0.25 || zeroOne.one != nil || zeroOne.zero != nil {
		t.Fatalf(`zeroOne.letter = %v, zeroOne.frequency = %v`, zeroOne.letter, zeroOne.frequency)
	}

	one := root.one
	if one.letter != "a" || one.frequency != 0.625 || one.one != nil || one.zero != nil {
		t.Fatalf(`one.letter = %v, one.frequency = %v`, one.letter, one.frequency)
	}
}

func TestFromMapping(t *testing.T) {
	mapping := make(map[string][]bool)
	mapping["a"] = []bool{true}
	mapping["b"] = []bool{false, true}
	mapping["c"] = []bool{false, false}

	root := fromMapping(mapping)
	if root.letter != "" || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`root.letter = %v, root.frequency = %v`, root.letter, root.frequency)
	}

	zero := root.zero
	if zero.letter != "" || zero.frequency != 0.0 || zero.one == nil || zero.zero == nil {
		t.Fatalf(`zero.letter = %v, zero.frequency = %v`, zero.letter, zero.frequency)
	}

	zeroZero := zero.zero
	if zeroZero.letter != "c" || zeroZero.frequency != 0.0 || zeroZero.one != nil || zeroZero.zero != nil {
		t.Fatalf(`zeroZero.letter = %v, zeroZero.frequency = %v`, zeroZero.letter, zeroZero.frequency)
	}

	zeroOne := zero.one
	if zeroOne.letter != "b" || zeroOne.frequency != 0.0 || zeroOne.one != nil || zeroOne.zero != nil {
		t.Fatalf(`zeroOne.letter = %v, zeroOne.frequency = %v`, zeroOne.letter, zeroOne.frequency)
	}

	one := root.one
	if one.letter != "a" || one.frequency != 0.0 || one.one != nil || one.zero != nil {
		t.Fatalf(`one.letter = %v, one.frequency = %v`, one.letter, one.frequency)
	}
}
