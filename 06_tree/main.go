package main

import (
	"fmt"

	"guochao.com/datastruct/06_tree/huffman"
)

func main() {
	ht := huffman.DefaultHuffmanTree()
	res := huffman.CreateHuffmanCode(ht, 7)
	for i, s := range res {
		fmt.Printf("%d: %s\n", i, s)
	}
}
