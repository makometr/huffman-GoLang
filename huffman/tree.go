package huffman

import "container/heap"

type haffmanBTNode struct {
	chars  []rune
	weight int
	left   *haffmanBTNode
	right  *haffmanBTNode
}

func mergeHuffmanBTNodes(lhs, rhs *haffmanBTNode) *haffmanBTNode {
	newRunes := make([]rune, 0, len(lhs.chars)+len(rhs.chars))
	newRunes = append(newRunes, lhs.chars...)
	newRunes = append(newRunes, rhs.chars...)
	return &haffmanBTNode{chars: newRunes, weight: lhs.weight + rhs.weight, left: lhs, right: rhs}
}

func (n haffmanBTNode) IsLeaf() bool {
	if n.left == nil && n.right == nil {
		return true
	}
	return false
}

func buildHuffmunTree(charFreq freqsTable) *haffmanBTNode {
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
