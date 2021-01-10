package huffman

import (
	"bytes"
	"encoding/gob"
	"log"
	"strings"

	"github.com/willf/bitset"
)

// Decompress decompresses the given data back into a string.
func Decompress(compressedContent []byte) string {
	exported := exportFormat{}

	decoder := gob.NewDecoder(bytes.NewReader(compressedContent))
	err := decoder.Decode(&exported)
	if err != nil {
		log.Fatal(err)
	}

	compressionResult := CompressionResult{*bitset.From(exported.Data), exported.Table, exported.Size}

	root := fromMapping(compressionResult.table)
	currentNode := root
	var text strings.Builder
	for i := uint(0); i < compressionResult.size; i++ {
		if compressionResult.data.Test(i) {
			currentNode = *currentNode.one
		} else {
			currentNode = *currentNode.zero
		}

		if currentNode.Leaf() {
			text.WriteRune(currentNode.letter)
			currentNode = root
		}
	}
	return text.String()
}
