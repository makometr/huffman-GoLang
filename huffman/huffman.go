package huffman

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
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
	return ""
	// TODO
	// return encodeByCodeTable(text, d.codes)
}
func (d AlgoData) DecodeText(text []byte) []byte {
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

// NewAlgoDataFromBytes returnes huffman full-intermediate data based on encoding string.
func NewAlgoDataFromBytes(text []byte) (*AlgoData, error) {
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

func EncodeBytes(byteText []byte) ([]byte, decodeTable, error) {
	res, table, err := Encode(byteText)
	if err != nil {
		return []byte{}, decodeTable{}, err
	}
	return []byte(res), table, nil
}

// Encode returnes encoded text by Huffman alogorithm.
func Encode(text []byte) ([]byte, decodeTable, error) {
	data, err := NewAlgoDataFromBytes(text)
	if err != nil {
		return []byte{}, decodeTable{}, err
	}
	size := countUsedBits(data.freqs, data.codes)
	// TODO in future size	/8 + 1
	return encodeWithTable(text, data.codes, size), ConvertEncodeTable(data.codes), nil
}

func encodeWithTable(text []byte, codeTable encodeTable, size int) []byte {
	// var builder strings.Builder
	var buffer bytes.Buffer
	for bytesRead := 0; bytesRead < len(text); {
		char, size := utf8.DecodeRune(text[bytesRead:])
		bytesRead += size
		// builder.WriteString(codeTable[char])
		buffer.Write([]byte(codeTable[char]))
	}
	// return []byte(builder.String())
	return buffer.Bytes()
}

func DecodeBytes(text []byte, table decodeTable) ([]byte, error) {
	result, err := DecodeString(text, table)
	return []byte(result), err
}

// DecodeString decodes input text by the decodeTable rules.
func DecodeString(text []byte, table decodeTable) ([]byte, error) {
	if err := checkUserDecodeTable(table); err != nil {
		return []byte{}, err
	}
	decodedText, err := decodeWithTable(text, table)
	if err != nil {
		return []byte{}, err
	}
	return decodedText, nil
}

func decodeWithTable(encodedText []byte, table decodeTable) ([]byte, error) {
	// TODO detect incorrect decoded text
	var resultBuilder strings.Builder
	curBeginIndex := 0

	for curEndIndex := 0; curEndIndex < len(encodedText); curEndIndex++ {
		currentBitSequence := encodedText[curBeginIndex : curEndIndex+1]
		if decodedChar, ok := table[string(currentBitSequence)]; ok {
			resultBuilder.WriteRune(decodedChar)
			curBeginIndex = curEndIndex + 1
		}
	}
	return []byte(resultBuilder.String()), nil
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

func checkUserFrequence(freqs freqsTable) error {
	for _, freq := range freqs {
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
		// Fix this
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

func countUsedBits(freq freqsTable, table encodeTable) (summ int) {
	for char, freq := range freq {
		summ += freq * len(table[char])
	}
	return
}

func printFreqs(freq freqsTable) {
	ordered := make([]struct {
		key   rune
		value int
	}, 0, len(freq))
	summ := 0
	for char, freq := range freq {
		ordered = append(ordered, struct {
			key   rune
			value int
		}{char, freq})
		summ += freq
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].value > ordered[j].value
	})
	for i := 0; i < len(ordered); i++ {
		fmt.Printf("%s (%d): %d\n", string(ordered[i].key), ordered[i].key, ordered[i].value)
	}
	fmt.Println("Unique chas: ", len(ordered))
	fmt.Println("Count of chars: ", summ)
}
