package main

import (
	"fmt"
	"huffmanApp/huffman"
)

func main() {
	var text string
	fmt.Scanf("%s\n", &text)
	msg, err := huffman.Encode(text)
	// data := huffman.AlgoData{}
	if err != nil {
		panic("srror")
	}
	// fmt.Printf("%d %d\n", len(encoder.charFreq), len(msg))
	// for char, code := range encoder.codes {
	// 	fmt.Printf("%c: %s\n", char, code)
	// }
	fmt.Println(msg)
}
