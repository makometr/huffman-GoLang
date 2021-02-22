package huffman

import (
	"reflect"
	"testing"
)

func Test_mergeHuffmanBTNodes(t *testing.T) {
	type args struct {
		lhs *haffmanBTNode
		rhs *haffmanBTNode
	}
	tests := []struct {
		name string
		args args
		want *haffmanBTNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeHuffmanBTNodes(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeHuffmanBTNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_haffmanBTNode_IsLeaf(t *testing.T) {
	testNode := haffmanBTNode{}
	tests := []struct {
		name string
		n    haffmanBTNode
		want bool
	}{
		{
			"Two sons.",
			haffmanBTNode{left: &testNode, right: &testNode},
			false,
		},
		{
			"One left son.",
			haffmanBTNode{left: &testNode},
			false,
		},
		{
			"One right son.",
			haffmanBTNode{right: &testNode},
			false,
		},
		{
			"No sons.",
			haffmanBTNode{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsLeaf(); got != tt.want {
				t.Errorf("haffmanBTNode.IsLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
