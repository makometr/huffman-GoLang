package huffman


type heapOfNodes []*haffmanBTNode

func (h heapOfNodes) Len() int {
	return len(h)
}
func (h heapOfNodes) Less(i, j int) bool {
	// if h[i].weight == h[j].weight {
	// 	return h[i].weight < h[j].weight
	// }
	return h[i].weight < h[j].weight
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
