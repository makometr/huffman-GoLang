package huffman

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type freqsTable map[rune]int
type encodeTable map[rune]string
type decodeTable map[string]rune

func ConvertEncodeTableToDecode(encode encodeTable) decodeTable {
	decode := make(decodeTable)
	for char, code := range encode {
		decode[code] = char
	}
	return decode
}

func ConvertDecodeTableToEncode(decode decodeTable) encodeTable {
	encode := make(encodeTable)
	for code, char := range decode {
		encode[char] = code
	}
	return encode
}

// AlgoData is
type AlgoData struct {
	freqs    freqsTable
	codes    encodeTable
	treeRoot *haffmanBTNode
}

func (d AlgoData) GetCharFrequences() freqsTable {
	return d.freqs
}
func (d AlgoData) GetCharCodes() encodeTable {
	return d.codes
}
func (d AlgoData) PrintTree() {
	fmt.Println("this is tree! todo")
}
func (d AlgoData) EncodeText(text string) string {
	return encodeByCodeTable(text, d.codes)
}
func (d AlgoData) DecodedText(text string) string {
	return decodeByDecodeTable(text, ConvertEncodeTableToDecode(d.codes))
}

// NewAlgoDataFromText returnes huffman full-intermediate data based on encoding string.
func NewAlgoDataFromText(text string) (*AlgoData, error) {
	if err := checkUserText(text); err != nil {
		return nil, err
	}
	data := newFilledData(countFreqs(text))
	return data, nil
}

// NewAlgoDataFromText returnes huffman full-intermediate data based on char frequence table.
func NewAlgoDataFromFrequence(freq freqsTable) (*AlgoData, error) {
	if err := checkUserFrequence(freq); err != nil {
		return nil, err
	}
	data := newFilledData(freq)
	return data, nil
}

// Encode returnes encoded text by Huffman alogorithm.
func Encode(text string) (string, error) {
	var result string
	if data, err := NewAlgoDataFromText(text); err != nil {
		return "", err
	} else {
		result = encodeByCodeTable(text, data.codes)
	}

	return result, nil
}

func Decode(text string, table decodeTable) (string, error) {
	// TODO check table
	return decodeByDecodeTable(text, table), nil
}

func decodeByDecodeTable(text string, table decodeTable) string {
	var resultBuilder strings.Builder
	curBeginIndex := 0
	for curEndIndex, _ := range text {
		currentBitSequence := text[curBeginIndex : curEndIndex+1]
		if decodedChar, ok := table[currentBitSequence]; ok {
			resultBuilder.WriteRune(decodedChar)
			curBeginIndex = curEndIndex + 1
		}
	}
	return resultBuilder.String()
}

func newAlgoData() *AlgoData {
	data := AlgoData{}
	data.freqs = freqsTable{}
	data.codes = encodeTable{}
	return &data
}

func newFilledData(freqs freqsTable) *AlgoData {
	data := newAlgoData()
	data.freqs = freqs
	data.treeRoot = buildHuffmunTree(data.freqs)
	data.codes = countOptimalCodes(data.treeRoot)
	return data
}

func (h AlgoData) print() {
	fmt.Println("Frequency:")
	for char, freq := range h.freqs {
		fmt.Printf("%c: %d\n", char, freq)
	}
	fmt.Printf("\nCodes:\n")
	for char, code := range h.codes {
		fmt.Printf("%c: %s\n", char, code)
	}
}

func countFreqs(text string) freqsTable {
	freqs := freqsTable{}
	for _, char := range text {
		freqs[char]++
	}
	return freqs
}

func countOptimalCodes(root *haffmanBTNode) encodeTable {
	codes := encodeTable{}
	generateCodesByTreeTraverse(root, codes)
	return codes
}

func generateCodesByTreeTraverse(root *haffmanBTNode, codes encodeTable) {
	if root.IsLeaf() {
		codes[root.chars[0]] = "0"
		return
	}
	var traverse func(rootNode *haffmanBTNode, prevCode string)
	traverse = func(rootNode *haffmanBTNode, prevCode string) {
		if rootNode.IsLeaf() {
			if len(rootNode.chars) != 1 {
				panic("Leaf has != 1 lenght of chars")
			}
			codes[rootNode.chars[0]] = prevCode
			return
		}
		traverse(rootNode.left, prevCode+"0")
		traverse(rootNode.right, prevCode+"1")
	}
	traverse(root, "")
}

func encodeByCodeTable(text string, codeTable encodeTable) string {
	var builder strings.Builder
	for _, char := range text {
		builder.WriteString(codeTable[char])
	}
	return builder.String()
}

func checkUserText(text string) error {
	for _, char := range text {
		if !unicode.IsGraphic(char) {
			return errors.New("Unicode graphic chars expected.")
		}
	}
	return nil
}

func checkUserFrequence(freqs freqsTable) error {
	for char, freq := range freqs {
		if !unicode.IsGraphic(char) {
			return errors.New("Unicode graphic chars expected.")
		}
		if freq <= 0 {
			return errors.New("Frequence must be positive integer.")
		}
	}
	return nil
}
