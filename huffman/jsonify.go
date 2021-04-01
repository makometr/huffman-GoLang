package huffman

import (
	"encoding/json"
	"fmt"
)

// TODO convert to binary code

type algoDataJSON struct {
	Encodedtext []byte
	DecodeTable decodeTable
}

func EncodeToJSON(text []byte) ([]byte, error) {
	encodedText, decodeTable, err := Encode(text)
	if err != nil {
		return []byte{}, err
	}
	bytes, err := json.Marshal(algoDataJSON{Encodedtext: encodedText, DecodeTable: decodeTable})
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func DecodeFromJSON(bytes []byte) ([]byte, error) {
	var algoData algoDataJSON
	if err := json.Unmarshal(bytes, &algoData); err != nil {
		return []byte{}, err
	}
	decodedText, err := DecodeString(algoData.Encodedtext, algoData.DecodeTable)
	if err != nil {
		return []byte{}, err
	}
	return decodedText, nil
}

func PrintEncodeStatisticsJSON(text []byte) {
	encodedBytes, err := EncodeToJSON(text)
	if err != nil {
		panic(err)
	}
	originalBytes, err := json.Marshal(struct{ Text []byte }{text})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Size of original text JSON: %d\n", len(originalBytes))
	fmt.Printf("Size of encoded  text JSON: %d\n", len(encodedBytes))
	fmt.Printf("Efficiency: %.2f%%\n", float64(len(originalBytes)-len(encodedBytes))/float64(len(originalBytes))*100)
}
