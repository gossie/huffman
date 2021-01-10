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
	if string(root.letter) != "a" || root.frequency != 1.0 || root.one != nil || root.zero != nil {
		t.Fatalf(`node.letter = %v, node.frequency = %v`, string(root.letter), root.frequency)
	}
}

func TestFromInputWithTwoDistinctLetters(t *testing.T) {
	root := fromInput("aaba")
	if root.letter != -1 || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`string(root.letter) = %v, root.frequency = %v`, string(root.letter), root.frequency)
	}

	zero := root.zero
	if string(zero.letter) != "b" || zero.frequency != 0.25 || zero.one != nil || zero.zero != nil {
		t.Fatalf(`string(zero.letter) = %v, zero.frequency = %v`, string(zero.letter), zero.frequency)
	}

	one := root.one
	if string(one.letter) != "a" || one.frequency != 0.75 || one.one != nil || one.zero != nil {
		t.Fatalf(`string(one.letter) = %v, one.frequency = %v`, string(one.letter), one.frequency)
	}
}

func TestFromInputWithThreeDistinctLetters(t *testing.T) {
	root := fromInput("aabaabca")
	if root.letter != -1 || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`string(root.letter) = %v, root.frequency = %v`, string(root.letter), root.frequency)
	}

	zero := root.zero
	if zero.letter != -1 || zero.frequency != 0.375 || zero.one == nil || zero.zero == nil {
		t.Fatalf(`string(zero.letter) = %v, zero.frequency = %v`, string(zero.letter), zero.frequency)
	}

	zeroZero := zero.zero
	if string(zero.zero.letter) != "c" || zeroZero.frequency != 0.125 || zeroZero.one != nil || zeroZero.zero != nil {
		t.Fatalf(`zerostring(zero.letter) = %v, zeroZero.frequency = %v`, string(zero.zero.letter), zeroZero.frequency)
	}

	zeroOne := zero.one
	if string(zero.one.letter) != "b" || zeroOne.frequency != 0.25 || zeroOne.one != nil || zeroOne.zero != nil {
		t.Fatalf(`zerostring(one.letter) = %v, zeroOne.frequency = %v`, string(zero.one.letter), zeroOne.frequency)
	}

	one := root.one
	if string(one.letter) != "a" || one.frequency != 0.625 || one.one != nil || one.zero != nil {
		t.Fatalf(`string(one.letter) = %v, one.frequency = %v`, string(one.letter), one.frequency)
	}
}

func TestFromInputMississippi(t *testing.T) {
	root := fromInput("Mississippi")
	if root.letter != -1 || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`string(root.letter) = %v, root.frequency = %v`, string(root.letter), root.frequency)
	}

	one := root.one
	if one.letter != -1 || one.one == nil || one.zero == nil {
		t.Fatalf(`one.letter = %v, one.frequency = %v`, one.letter, one.frequency)
	}

	oneOne := one.one
	if string(oneOne.letter) != "i" || oneOne.one != nil || oneOne.zero != nil {
		t.Fatalf(`oneOne.letter = %v, oneOne.frequency = %v`, string(oneOne.letter), oneOne.frequency)
	}

	oneZero := one.zero
	if oneZero.letter != -1 || oneZero.one == nil || oneZero.zero == nil {
		t.Fatalf(`oneZero.letter = %v, oneZero.frequency = %v`, oneZero.letter, oneZero.frequency)
	}

	oneZeroOne := oneZero.one
	if string(oneZeroOne.letter) != "p" || oneZeroOne.one != nil || oneZeroOne.zero != nil {
		t.Fatalf(`oneZeroOne.letter = %v, oneZeroOne.frequency = %v`, string(oneZeroOne.letter), oneZeroOne.frequency)
	}

	oneZeroZero := oneZero.zero
	if string(oneZeroZero.letter) != "M" || oneZeroZero.one != nil || oneZeroZero.zero != nil {
		t.Fatalf(`oneZeroZero.letter = %v, oneZeroZero.frequency = %v`, string(oneZeroZero.letter), oneZeroZero.frequency)
	}

	zero := root.zero
	if string(zero.letter) != "s" || zero.one != nil || zero.zero != nil {
		t.Fatalf(`zero.letter = %v, zero.frequency = %v`, zero.letter, zero.frequency)
	}

	// zeroZero := zero.zero
	// if string(zero.zero.letter) != "c" || zeroZero.frequency != 0.125 || zeroZero.one != nil || zeroZero.zero != nil {
	// 	t.Fatalf(`zerostring(zero.letter) = %v, zeroZero.frequency = %v`, string(zero.zero.letter), zeroZero.frequency)
	// }

	// zeroOne := zero.one
	// if string(zero.one.letter) != "b" || zeroOne.frequency != 0.25 || zeroOne.one != nil || zeroOne.zero != nil {
	// 	t.Fatalf(`zerostring(one.letter) = %v, zeroOne.frequency = %v`, string(zero.one.letter), zeroOne.frequency)
	// }
}

func TestFromMapping(t *testing.T) {
	mapping := make(map[rune][]bool)
	mapping[stringToRune("a")] = []bool{true}
	mapping[stringToRune("b")] = []bool{false, true}
	mapping[stringToRune("c")] = []bool{false, false}

	root := fromMapping(mapping)
	if root.letter != -1 || root.frequency != 1.0 || root.one == nil || root.zero == nil {
		t.Fatalf(`string(root.letter) = %v, root.frequency = %v`, string(root.letter), root.frequency)
	}

	zero := root.zero
	if zero.letter != -1 || zero.frequency != 0.0 || zero.one == nil || zero.zero == nil {
		t.Fatalf(`string(zero.letter) = %v, zero.frequency = %v`, string(zero.letter), zero.frequency)
	}

	zeroZero := zero.zero
	if string(zero.zero.letter) != "c" || zeroZero.frequency != 0.0 || zeroZero.one != nil || zeroZero.zero != nil {
		t.Fatalf(`zerostring(zero.letter) = %v, zeroZero.frequency = %v`, string(zero.zero.letter), zeroZero.frequency)
	}

	zeroOne := zero.one
	if string(zero.one.letter) != "b" || zeroOne.frequency != 0.0 || zeroOne.one != nil || zeroOne.zero != nil {
		t.Fatalf(`zerostring(one.letter) = %v, zeroOne.frequency = %v`, string(zero.one.letter), zeroOne.frequency)
	}

	one := root.one
	if string(one.letter) != "a" || one.frequency != 0.0 || one.one != nil || one.zero != nil {
		t.Fatalf(`string(one.letter) = %v, one.frequency = %v`, string(one.letter), one.frequency)
	}
}

func stringToRune(s string) rune {
	return []rune(s)[0]
}
