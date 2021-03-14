package huffman

import (
	"container/heap"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// IntermediateData is
type IntermediateData struct {
	charFreq  map[rune]int
	codeTable map[rune]string
	treeRoot  *haffmanBTNode
}

func (d IntermediateData) getCharFrequences() map[rune]int {
	return d.charFreq
}
func (d IntermediateData) getCharCodes() map[rune]string {
	return d.codeTable
}
func (d IntermediateData) PrintTree() {
	fmt.Println("this is tree! todo")
}

// Encode method counts encodes text by Huffman algorithm
func Encode(text string) (string, error) {
	var optimal = IntermediateData{}
	if !isTextValid(text) {
		return "", errors.New("No-letter char in text.")
	}
	optimal.charFreq = countFreq(text)
	optimal.treeRoot = buildHuffmunTree(optimal.charFreq)
	optimal.codeTable = countOptimalCodes(optimal.treeRoot)
	newText := encodeByCodeTable(text, optimal.codeTable)

	return newText, nil
}

func (h IntermediateData) print() {
	fmt.Println("Frequency:")
	for char, freq := range h.charFreq {
		fmt.Printf("%c: %d\n", char, freq)
	}
	fmt.Printf("\nCodes:\n")
	for char, code := range h.codeTable {
		fmt.Printf("%c: %s\n", char, code)
	}
}

func isTextValid(text string) bool {
	for _, char := range text {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func countFreq(text string) map[rune]int {
	freqs := map[rune]int{}
	for _, char := range text {
		freqs[char]++
	}
	return freqs
}

func countOptimalCodes(root *haffmanBTNode) map[rune]string {
	codes := map[rune]string{}
	generateCodesByTreeTraverse(root, codes)
	return codes
}

func buildHuffmunTree(charFreq map[rune]int) *haffmanBTNode {
	nodes := make(heapOfNodes, 0, len(charFreq))
	for char, freq := range charFreq {
		nodes = append(nodes, &haffmanBTNode{chars: []rune{char}, weight: freq})
	}
	heap.Init(&nodes)

	for len(nodes) > 1 {
		leftNode := heap.Pop(&nodes).(*haffmanBTNode)
		rightNode := heap.Pop(&nodes).(*haffmanBTNode)
		newNode := mergeHuffmanBTNodes(leftNode, rightNode)
		heap.Push(&nodes, newNode)
	}
	return heap.Pop(&nodes).(*haffmanBTNode)
}

func generateCodesByTreeTraverse(root *haffmanBTNode, codes map[rune]string) {
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

func encodeByCodeTable(text string, codeTable map[rune]string) string {
	var builder strings.Builder
	for _, char := range text {
		builder.WriteString(codeTable[char])
	}
	return builder.String()
}
