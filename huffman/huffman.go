package huffman

import (
	"container/heap"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// OptCode is
type OptCode struct {
	charFreq map[rune]int
	codes    map[rune]string
}

// NewHuffmanOptCode is constrcutor
func NewHuffmanOptCode() *OptCode {
	var obj OptCode
	obj.charFreq = make(map[rune]int)
	obj.codes = make(map[rune]string)
	return &obj
}

// Encode method counts encodes text by Huffman algorithm
func (h *OptCode) Encode(text string) (string, error) {
	if !isTextValid(text) {
		return "", errors.New("No-letter char in text")
	}
	h.countFreq(text)
	h.countOptimalCodes()
	// h.print()
	newText := h.encodeText(text)

	return newText, nil
}

func (h OptCode) print() {
	fmt.Println("Frequency:")
	for char, freq := range h.charFreq {
		fmt.Printf("%c: %d\n", char, freq)
	}
	fmt.Printf("\nCodes:\n")
	for char, code := range h.codes {
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

func (h *OptCode) countFreq(text string) {
	for _, char := range text {
		h.charFreq[char]++
	}
}

func (h *OptCode) countOptimalCodes() {
	root := h.buildHuffmunTree()
	h.generateCodesByTreeTraverse(root)
}

func (h *OptCode) buildHuffmunTree() *haffmanBTNode {
	nodes := make(heapOfNodes, 0, len(h.charFreq))
	for char, freq := range h.charFreq {
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

func (h *OptCode) generateCodesByTreeTraverse(root *haffmanBTNode) {
	if root.IsLeaf() {
		h.codes[root.chars[0]] = "0"
		return
	}
	var traverse func(rootNode *haffmanBTNode, prevCode string)
	traverse = func(rootNode *haffmanBTNode, prevCode string) {
		if rootNode.IsLeaf() {
			if len(rootNode.chars) != 1 {
				panic("Leaf has != 1 lenght of chars")
			}
			h.codes[rootNode.chars[0]] = prevCode
			return
		}
		traverse(rootNode.left, prevCode+"0")
		traverse(rootNode.right, prevCode+"1")
	}
	traverse(root, "")
}

func (h *OptCode) encodeText(text string) string {
	var builder strings.Builder
	for _, char := range text {
		builder.WriteString(h.codes[char])
	}
	return builder.String()
}
