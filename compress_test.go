package huffman

import (
	"testing"
)

func TestCompress(t *testing.T) {
	compressed := CompressConcurrent("aabaabca")
	if compressed.Size() != 11 {
		t.Fatalf("compressed = %v", compressed)
	}
}

func TestCompressions(t *testing.T) {
	compressed1 := CompressSingle("aabaabca")
	compressed2 := CompressConcurrent("aabaabca")
	if compressed1.Size() != compressed2.Size() {
		t.Fatalf("%v != %v", compressed1, compressed2)
	}
}

func TestCompressAndDecompress(t *testing.T) {
	compressed := CompressConcurrent("aabaabca")
	decompressed := Decompress(compressed)
	if decompressed != "aabaabca" {
		t.Fatalf("decompressed = %v", decompressed)
	}
}
