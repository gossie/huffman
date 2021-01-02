package huffman

import (
	"testing"
)

func TestCreateTreeWithOnlyOneLetter(t *testing.T) {
	root := createTree("a")
	if root.one != nil || root.zero != nil {
		t.Fatal("The node should not have children")
	}
}

func TestCreateTreeWithOneDistinctLetter(t *testing.T) {
	root := createTree("aaa")
	if root.letter != "a" || root.frequency != 1.0 || root.one != nil || root.zero != nil {
		t.Fatalf(`node.letter = %v, node.frequency = %v`, root.letter, root.frequency)
	}
}

func TestCreateTreeWith(t *testing.T) {
	root := createTree("aaba")
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
