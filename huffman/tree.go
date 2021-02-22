package huffman

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
