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

func TestCompressAndDecompressWithOnlyOneLetter(t *testing.T) {
	compressed := Compress("aaaaaaaa")
	decompressed := Decompress(compressed.Bytes())
	if decompressed != "aaaaaaaa" {
		t.Fatalf("decompressed = %v", decompressed)
	}
}
