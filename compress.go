package huffman

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/willf/bitset"
)

// CompressionResult holds the compressed data and huffman tree.
type CompressionResult struct {
	data  bitset.BitSet
	table map[string][]bool
	size  uint
}

type exportFormat struct {
	Data  []uint64
	Table map[string][]bool
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
	letterChannel := make(chan string, 100000)
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
	root := fromInput(input)
	mapping := letterCodeMapping(&root)
	var data bitset.BitSet

	var index uint = 0
	for _, letter := range input {
		for _, bit := range mapping[string(letter)] {
			if bit {
				data.Set(index)
			}
			index++
		}
	}
	return CompressionResult{data, mapping, index}
}

func readLetter(c chan string, input string) {
	for _, letter := range input {
		c <- string(letter)
	}
	close(c)
}

func mapLetter(in chan string, out chan []bool, mapping map[string][]bool) {
	for letter := range in {
		out <- mapping[letter]
	}
	close(out)
}

func letterCodeMapping(root *node) map[string][]bool {
	mapping := make(map[string][]bool)
	toBeNamed(root, make([]bool, 0), mapping)
	return mapping
}

func toBeNamed(node *node, code []bool, mapping map[string][]bool) {
	if node.letter != "" {
		mapping[node.letter] = code
	} else {
		code1 := make([]bool, len(code))
		code0 := make([]bool, len(code))

		copy(code1, code)
		copy(code0, code)

		toBeNamed(node.one, append(code1, true), mapping)
		toBeNamed(node.zero, append(code0, false), mapping)
	}
}
