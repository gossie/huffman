package huffman

import "strings"

// Decompress decompresses the given data back into a string.
func Decompress(compressionResult CompressionResult) string {
	root := fromMapping(compressionResult.table)
	currentNode := root
	var text strings.Builder
	for i := uint(0); i < compressionResult.Size(); i++ {
		if compressionResult.data.Test(i) {
			currentNode = *currentNode.one
		} else {
			currentNode = *currentNode.zero
		}

		if currentNode.letter != "" {
			text.WriteString(currentNode.letter)
			currentNode = root
		}
	}
	return text.String()
}
