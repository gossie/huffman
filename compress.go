package huffman

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// CompressionResult holds the compressed data and huffman tree.
type CompressionResult struct {
	data  BitSet
	table map[rune][]byte
	size  uint
}

type exportFormat struct {
	Data  []byte
	Table map[rune][]byte
	Size  uint
}

// Bytes returns a byte array representation.
func (cr *CompressionResult) Bytes() []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(exportFormat{cr.data.Bytes(), cr.table, cr.size})
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

// Compress takes a string and compresses it using Huffman code.
func Compress(input string) CompressionResult {
	startTree := time.Now()
	root := fromInput(input)
	endTree := time.Now()
	log.Println("creating tree: ", endTree.UnixNano()-startTree.UnixNano())

	startMapping := time.Now()
	mapping := letterCodeMapping(&root)
	endMapping := time.Now()
	log.Println("creating mapping table: ", endMapping.UnixNano()-startMapping.UnixNano())

	var data BitSet

	var index uint = 0
	for _, letter := range input {
		code := mapping[letter]
		numberOfBits := code[len(mapping[letter])-1]
		for _, bits := range code[:len(code)-1] {
			if numberOfBits > 0 {
				setBit(&data, bits, 1, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 2, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 4, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 8, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 16, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 32, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 64, index)
				index++
				numberOfBits--
			}

			if numberOfBits > 0 {
				setBit(&data, bits, 1<<7, index)
				index++
				numberOfBits--
			}
		}
	}
	// fmt.Println("compressed: ", data)
	// fmt.Println("mapping: ", mapping)
	return CompressionResult{data, mapping, index}
}

func setBit(data *BitSet, bits byte, mask byte, index uint) {
	if (bits & mask) != 0 {
		data.Set(index)
	}
}

func letterCodeMapping(root *node) map[rune][]byte {
	mapping := make(map[rune][]byte)
	toBeNamed(root, make([]byte, 0), &mapping, 0)
	return mapping
}

func toBeNamed(node *node, code []byte, mapping *map[rune][]byte, numberOfBits byte) {
	if node.Leaf() {
		if len(code) > 0 {
			(*mapping)[node.letter] = append(code, numberOfBits)
		} else {
			(*mapping)[node.letter] = append(code, 1, 1)
		}
	} else {
		code1 := make([]byte, len(code))
		code0 := make([]byte, len(code))

		copy(code1, code)
		copy(code0, code)

		bitIndex := numberOfBits % 8
		if bitIndex == 0 {
			code1 = append(code1, 1)
			code0 = append(code0, 0)
		} else {
			code1[len(code1)-1] |= (1 << bitIndex)
		}

		toBeNamed(node.one, code1, mapping, numberOfBits+1)
		toBeNamed(node.zero, code0, mapping, numberOfBits+1)
	}
}
