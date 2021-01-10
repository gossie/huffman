package huffman

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"

	"github.com/willf/bitset"
)

// CompressionResult holds the compressed data and huffman tree.
type CompressionResult struct {
	data  bitset.BitSet
	table map[rune][]bool
	size  uint
}

type exportFormat struct {
	Data  []uint64
	Table map[rune][]bool
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

func compressConcurrent(input string) CompressionResult {
	root := fromInput(input)
	mapping := letterCodeMapping(&root)
	var data bitset.BitSet
	letterChannel := make(chan rune, 100000)
	mappingChannel := make(chan []bool, 100000)

	go readLetter(letterChannel, input)
	go mapLetter(letterChannel, mappingChannel, mapping)

	var index uint = 0
	for code := range mappingChannel {
		for _, bit := range code {
			if bit {
				data.Set(index)
			}
			index++
		}
	}

	return CompressionResult{data, mapping, index}
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

	var data bitset.BitSet

	var index uint = 0
	for _, letter := range input {
		for _, bit := range mapping[letter] {
			if bit {
				data.Set(index)
			}
			index++
		}
	}
	return CompressionResult{data, mapping, index}
}

func readLetter(c chan rune, input string) {
	for _, letter := range input {
		c <- letter
	}
	close(c)
}

func mapLetter(in chan rune, out chan []bool, mapping map[rune][]bool) {
	for letter := range in {
		out <- mapping[letter]
	}
	close(out)
}

func letterCodeMapping(root *node) map[rune][]bool {
	mapping := make(map[rune][]bool)
	toBeNamed(root, make([]bool, 0), &mapping)
	return mapping
}

func toBeNamed(node *node, code []bool, mapping *map[rune][]bool) {
	if node.letter != -1 {
		if len(code) > 0 {
			(*mapping)[node.letter] = code
		} else {
			(*mapping)[node.letter] = append(code, true)
		}
	} else {
		code1 := make([]bool, len(code))
		code0 := make([]bool, len(code))

		copy(code1, code)
		copy(code0, code)

		toBeNamed(node.one, append(code1, true), mapping)
		toBeNamed(node.zero, append(code0, false), mapping)
	}
}
