package main

import (
	"fmt"
	"huffmanApp/huffman"
)

func main() {
	var text string
	fmt.Scanf("%s\n", &text)

	data, _ := huffman.NewAlgoDataFromText(text)
	encodedText := data.EncodeText(text)
	// decodeTable := huffman.ConvertEncodeTableToDecode(data.GetCharCodes())
	decodedText := data.DecodedText(encodedText)
	fmt.Println(encodedText)
	fmt.Println(decodedText)
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
