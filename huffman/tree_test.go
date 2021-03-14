package huffman

import (
	"reflect"
	"testing"
)

func Test_mergeHuffmanBTNodes(t *testing.T) {
	nilNode1 := haffmanBTNode{}
	nilNode2 := haffmanBTNode{}
	nilNodeResult := haffmanBTNode{left: &nilNode1, right: &nilNode2, chars: []rune{}}

	leaf11 := haffmanBTNode{weight: 10, chars: []rune{'a', 'b', 'c'}}
	leaf12 := haffmanBTNode{weight: 15, chars: []rune{'d', 'e'}}
	node1 := haffmanBTNode{weight: 25, left: &leaf11, right: &leaf12, chars: []rune{'a', 'b', 'c', 'd', 'e'}}

	leaf21 := haffmanBTNode{weight: 20, chars: []rune{'f', 'e', 'p'}}
	leaf22 := haffmanBTNode{weight: 20, chars: []rune{'a', 'v', 's'}}
	node2 := haffmanBTNode{weight: 40, left: &leaf21, right: &leaf22, chars: []rune{'f', 'e', 'p', 'a', 'v', 's'}}
	type args struct {
		lhs *haffmanBTNode
		rhs *haffmanBTNode
	}
	tests := []struct {
		name string
		args args
		want *haffmanBTNode
	}{
		{
			"Nil nodes.",
			args{&nilNode1, &nilNode2},
			&nilNodeResult,
		},
		{
			"Full check (different weight).",
			args{&leaf11, &leaf12},
			&node1,
		},
		{
			"Full check (same weight, incorrect order).",
			args{&leaf21, &leaf22},
			&node2,
		},
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
