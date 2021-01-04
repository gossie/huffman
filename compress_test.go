package huffman

import (
	"testing"
)

func TestCompressAndDecompress(t *testing.T) {
	compressed := Compress("aabaabca")
	decompressed := Decompress(compressed.Bytes())
	if decompressed != "aabaabca" {
		t.Fatalf("decompressed = %v", decompressed)
	}
}
