package huffman

import (
	"log"
)

type heapOfNodes []*haffmanBTNode

func (h heapOfNodes) Len() int {
	return len(h)
}
func (h heapOfNodes) Less(i, j int) bool {
	if len(h[i].chars) == 0 || len(h[j].chars) == 0 {
		log.Println("All compared nodes must be non-empty.")
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	if len(h[i].chars) != len(h[j].chars) {
		return len(h[i].chars) < len(h[j].chars)
	}
	if h[i].chars[0] == h[j].chars[0] {
		log.Println("Same characters are not expected.")
	}
	return h[i].chars[0] < h[j].chars[0]
}

func (h heapOfNodes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *heapOfNodes) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *heapOfNodes) Push(elem interface{}) {
	*h = append(*h, elem.(*haffmanBTNode))
}
