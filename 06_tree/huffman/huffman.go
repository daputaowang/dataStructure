package huffman

import (
	"fmt"
)

type HuffmanTreeNode struct {
	weight           float64
	parent, lch, rch int
}

func DefaultHuffmanTree() []HuffmanTreeNode {
	return []HuffmanTreeNode{
		{},
		{0.4, 13, 0, 0},
		{0.3, 12, 0, 0},
		{0.15, 11, 0, 0},
		{0.05, 9, 0, 0},
		{0.04, 9, 0, 0},
		{0.03, 8, 0, 0},
		{0.03, 8, 0, 0},
		{0.06, 10, 7, 6},
		{0.09, 10, 5, 4},
		{0.15, 11, 8, 9},
		{0.3, 12, 10, 3},
		{0.6, 13, 11, 2},
		{1, 0, 1, 12},
	}
}

func CreateHuffmanCode(ht []HuffmanTreeNode, n int) []string {
	hc := make([]string, n+1)
	for i := 1; i <= n; i++ {
		f := ht[i].parent
		c := i
		cd := ""
		for f != 0 {
			if ht[f].lch == c {
				cd = fmt.Sprintf("%s%s", "0", cd)
			} else {
				cd = fmt.Sprintf("%s%s", "1", cd)
			}
			c = f
			f = ht[f].parent
		}
		hc[i] = cd
		cd = ""
	}
	return hc
}
