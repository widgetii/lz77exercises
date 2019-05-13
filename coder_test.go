package lz77

import (
	"bytes"
	"testing"
)

func TestSimple(t *testing.T) {
	input := []byte("abcbaadbcbaadacdc")
	//input := []byte("aabcbbabc")
	if !bytes.Equal(Lz77Encode(input), input) {
		t.Errorf("Lz77Encode(%q) != %q", input, input)
	}
}
