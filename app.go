package main

import (
	"fmt"
	"huffmanApp/huffman"
)

func main() {
	var text string
	fmt.Scanf("%s\n", &text)

	func() {
		// Example of using module with full intermediate data object.
		data, _ := huffman.NewAlgoDataFromText(text)
		encodedText := data.EncodeText(text)
		decodedText := data.DecodeText(encodedText)
		fmt.Println(encodedText)
		fmt.Println(decodedText)
		data.PrintStatistics()
		huffman.PrintEncodeStatisticsJSON(text)
	}()

	// func() {
	// 	jsonObj, err := huffman.EncodeToJSON(text)
	// 	if err != nil {
	// 		panic(err)
	// 	} else {
	// 		fmt.Println(string(jsonObj))
	// 	}

	// 	decodedText, err := huffman.DecodeFromJSON(jsonObj)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(decodedText)
	// 	fmt.Println(decodedText == text)
	// }()

	// func() {
	// 	// Example of using module without proxy-object.
	// 	encodedText, decodeTable, err := huffman.Encode(text)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if decodedText, err := huffman.Decode(encodedText, decodeTable); err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(decodedText)
	// 	}
	// }()
}
