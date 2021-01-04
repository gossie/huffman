package huffman

import (
	"testing"
)

func TestCompress(t *testing.T) {
	compressed := CompressConcurrent("aabaabca")
	if compressed.size != 11 {
		t.Fatalf("compressed = %v", compressed)
	}
}

func TestCompressions(t *testing.T) {
	compressed1 := CompressSingle("aabaabca")
	compressed2 := CompressConcurrent("aabaabca")
	if compressed1.size != compressed2.size {
		t.Fatalf("%v != %v", compressed1, compressed2)
	}
}

func TestCompressAndDecompress(t *testing.T) {
	compressed := CompressConcurrent("aabaabca")
	decompressed := DecompressSingle(compressed)
	if decompressed != "aabaabca" {
		t.Fatalf("decompressed = %v", decompressed)
	}
}
