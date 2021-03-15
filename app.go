package main

import (
	"fmt"
	"huffmanApp/huffman"
)

func main() {
	var text string
	fmt.Scanf("%s\n", &text)

	data, _ := huffman.NewAlgoDataFromText(text)
	encodedText, decodeTable := data.EncodeText(text)
	decodedText := data.DecodeText(encodedText)
	fmt.Println(encodedText)
	fmt.Println(decodedText)

	if decodedText, err := huffman.Decode(encodedText, decodeTable); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(decodedText)
	}

	// msg, err := huffman.Encode(text)
	// if err != nil {
	// 	panic("srror")
	// }
	// fmt.Printf("%d %d\n", len(encoder.charFreq), len(msg))
	// for char, code := range encoder.codes {
	// 	fmt.Printf("%c: %s\n", char, code)
	// }
	// fmt.Println(msg)
}
