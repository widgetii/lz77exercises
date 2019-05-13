package lz77

import (
	"fmt"
)

const (
	windowSize = 0xFF
	aheadSize  = 0xF
)

func Lz77Encode(data []byte) (out []byte) {

	for windowCenter := 0; windowCenter < len(data); {
		// search inside window from end to start
		maxSeq := 0
		var maxI int
		// TODO: check 2nd cond
		for i := windowCenter - 1; i >= 0 && i > windowCenter-windowSize; i-- {
			if data[i] == data[windowCenter] {
				// check how much len of sequence
				seqLen := 1
				for j := i + 1; windowCenter+seqLen < len(data) && data[windowCenter+seqLen] == data[j]; j++ {
					seqLen++
				}
				if seqLen >= maxSeq {
					maxSeq = seqLen
					maxI = i
				}
			}
		}
		if maxSeq != 0 {
			ch := ' '
			if (windowCenter+maxSeq+1) < len(data) {
				ch = rune(data[windowCenter+maxSeq+1])
			}
			fmt.Printf("[%d,%d]%c\n", windowCenter-maxI, maxSeq, ch)
			windowCenter+=maxSeq+1
		} else {
			fmt.Printf("[0,0]%c\n", data[windowCenter])
			windowCenter++
		}
	}

	// while data is available, process it
	for _, v := range data {
		out = append(out, v)
	}
	return
}

func main() {
	fmt.Println("vim-go")
}
