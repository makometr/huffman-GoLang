package main

import (
	"bytes"
	"fmt"
	"huffmanApp/huffman"
	"io/ioutil"
)

func main() {
	text, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	encodedText, table, err := huffman.EncodeBytes(text)
	if err != nil {
		panic(err)
	}
	// fmt.Println(encodedText)
	fmt.Println(len(encodedText))
	result, err := huffman.DecodeBytes(encodedText, table)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(result))
	if err := ioutil.WriteFile("result.txt", result, 0777); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Size original: ", len(text))
	fmt.Println("Size after encode-decode: ", len(result))
	fmt.Println(bytes.Compare(text, result))

	// fmt.Println(string(text))

	// func() {
	// 	// Example of using module with full intermediate data object.
	// 	data, _ := huffman.NewAlgoDataFromText(text)
	// 	encodedText := data.EncodeText(text)
	// 	decodedText := data.DecodeText(encodedText)
	// 	fmt.Println(encodedText)
	// 	fmt.Println(decodedText)
	// 	data.PrintStatistics()
	// 	huffman.PrintEncodeStatisticsJSON(text)
	// }()

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
