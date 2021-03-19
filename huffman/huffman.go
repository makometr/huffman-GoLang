package huffman

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type freqsTable map[rune]int
type encodeTable map[rune]string
type decodeTable map[string]rune

func ConvertEncodeTable(encode encodeTable) decodeTable {
	decode := make(decodeTable)
	for char, code := range encode {
		decode[code] = char
	}
	return decode
}

func ConvertDecodeTable(decode decodeTable) encodeTable {
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
func (d AlgoData) GetEncodeTable() encodeTable {
	return d.codes
}
func (d AlgoData) GetDecodeTable() decodeTable {
	return ConvertEncodeTable(d.codes)
}
func (d AlgoData) PrintTree() {
	fmt.Println("this is tree! todo")
}
func (d AlgoData) EncodeText(text string) string {
	return encodeByCodeTable(text, d.codes)
}
func (d AlgoData) DecodeText(text string) string {
	// TODO is error expected ???
	decodedtext, _ := decodeWithTable(text, ConvertEncodeTable(d.codes))
	return decodedtext
}
func (d AlgoData) PrintStatistics() {
	var fullEncodedSize, fullOriginalSize, sizeOfDecodeTable int
	for char, freq := range d.freqs {
		sizeOriginal := len(string(char))
		encodedSize := len(d.codes[char])
		fullOriginalSize += sizeOriginal * freq
		fullEncodedSize += encodedSize * freq
		sizeOfDecodeTable += (sizeOriginal + 1 + encodedSize + 1)
	}
	fullEncodedSize = fullEncodedSize/8 + 1
	realEncodedSize := fullEncodedSize + sizeOfDecodeTable

	fmt.Printf("Size of non-encoded text: %d bytes.\n", fullOriginalSize)
	fmt.Printf("Size of encoded text + size of decode-table: %d+%d=%d bytes.\n", fullEncodedSize, sizeOfDecodeTable, realEncodedSize)
	fmt.Printf("Efficiency: %.2f%%\n", float64(fullOriginalSize-realEncodedSize)/float64(fullOriginalSize)*100)
}

// NewAlgoDataFromText returnes huffman full-intermediate data based on encoding string.
func NewAlgoDataFromText(text string) (*AlgoData, error) {
	if err := checkUserText(text); err != nil {
		return nil, err
	}
	data := newFilledData(countFreqs([]byte(text)))
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

func EncodeBytes(byteText []byte) ([]byte, decodeTable, error) {
	// TODO do EncodeString a wrapper over EncodeBytes
	res, table, err := EncodeString(string(byteText))
	if err != nil {
		return []byte{}, decodeTable{}, err
	}
	return []byte(res), table, nil
}

// EncodeString returnes encoded text by Huffman alogorithm.
func EncodeString(text string) (string, decodeTable, error) {
	data, err := NewAlgoDataFromText(text)
	if err != nil {
		return "", decodeTable{}, err
	}
	return encodeByCodeTable(text, data.codes), ConvertEncodeTable(data.codes), nil
}

func encodeByCodeTable(text string, codeTable encodeTable) string {
	var builder strings.Builder
	for _, char := range text {
		builder.WriteString(codeTable[char])
	}
	return builder.String()
}

func DecodeBytes(text []byte, table decodeTable) ([]byte, error) {
	result, err := DecodeString(string(text), table)
	return []byte(result), err
}

// DecodeString decodes input text by the decodeTable rules.
func DecodeString(text string, table decodeTable) (string, error) {
	if err := checkUserDecodeTable(table); err != nil {
		return "", err
	}
	decodedText, err := decodeWithTable(text, table)
	if err != nil {
		return "", err
	}
	return decodedText, nil
}

func decodeWithTable(encodedText string, table decodeTable) (string, error) {
	// TODO detect incorrect decoded text
	var resultBuilder strings.Builder
	curBeginIndex := 0
	for curEndIndex, _ := range encodedText {
		currentBitSequence := encodedText[curBeginIndex : curEndIndex+1]
		if decodedChar, ok := table[currentBitSequence]; ok {
			resultBuilder.WriteRune(decodedChar)
			curBeginIndex = curEndIndex + 1
		}
	}
	return resultBuilder.String(), nil
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

func (h AlgoData) Print() {
	fmt.Println("Frequency:")
	for char, freq := range h.freqs {
		fmt.Printf("%c: %d\n", char, freq)
	}
	fmt.Printf("\nCodes:\n")
	for char, code := range h.codes {
		fmt.Printf("%c: %s\n", char, code)
	}
}

func countFreqs(text []byte) freqsTable {
	freqs := freqsTable{}
	for bytesRead := 0; bytesRead < len(text); {
		char, size := utf8.DecodeRune(text[bytesRead:])
		bytesRead += size
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

func checkUserText(text string) error {
	// TODO whats check?)
	return nil
}

func checkUserFrequence(freqs freqsTable) error {
	for char, freq := range freqs {
		if !unicode.IsGraphic(char) {
			return errors.New("unicode graphic chars expected")
		}
		if freq <= 0 {
			return errors.New("frequence must be positive integer")
		}
	}
	return nil
}

func checkUserDecodeTable(table decodeTable) error {
	uniqueChars := make(map[rune]bool)
	for code, char := range table {
		if _, ok := uniqueChars[char]; ok {
			return errors.New("unique chars encoded expected")
		}
		samePrefixCounter := 0
		for checkedCode, _ := range table {
			if strings.HasPrefix(code, checkedCode) {
				samePrefixCounter++
			}
		}
		if samePrefixCounter > 1 {
			return errors.New(fmt.Sprintln("table contains codes with same prefixes"))
		}
	}
	return nil
}
