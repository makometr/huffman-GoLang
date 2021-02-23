package huffman

import (
	"reflect"
	"testing"
)

func Test_heapOfNodes_Len(t *testing.T) {
	tests := []struct {
		name string
		h    heapOfNodes
		want int
	}{
		{
			"No elements.",
			heapOfNodes{},
			0,
		},
		{
			"One element.",
			heapOfNodes{&haffmanBTNode{}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("heapOfNodes.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapOfNodes_Less(t *testing.T) {
	charsD := []rune{'D'}
	charsAB := []rune{'A', 'B'}
	charsFS := []rune{'F', 'S'}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    heapOfNodes
		args args
		want bool
	}{
		{
			"Different-weight nodes. Different chars, diffrent length.",
			heapOfNodes{&haffmanBTNode{weight: 0, chars: charsAB}, &haffmanBTNode{weight: 10, chars: charsD}},
			args{0, 1},
			true,
		},
		{
			"Different-weight nodes. Different chars, same length.",
			heapOfNodes{&haffmanBTNode{weight: 0, chars: charsAB}, &haffmanBTNode{weight: 10, chars: charsFS}},
			args{0, 1},
			true,
		},
		{
			"Different-weight nodes. Same chars-1.",
			heapOfNodes{&haffmanBTNode{weight: 0, chars: charsD}, &haffmanBTNode{weight: 10, chars: charsD}},
			args{0, 1},
			true,
		},
		{
			"Different-weight nodes. Same chars-2.",
			heapOfNodes{&haffmanBTNode{weight: 0, chars: charsAB}, &haffmanBTNode{weight: 10, chars: charsAB}},
			args{0, 1},
			true,
		},
		{
			"Same weight nodes. Different chars, diffrent length.",
			heapOfNodes{&haffmanBTNode{weight: 10, chars: charsD}, &haffmanBTNode{weight: 10, chars: charsAB}},
			args{0, 1},
			true,
		},
		{
			"Same weight nodes. Different chars, same length.",
			heapOfNodes{&haffmanBTNode{weight: 10, chars: charsAB}, &haffmanBTNode{weight: 10, chars: charsFS}},
			args{0, 1},
			true,
		},
		{
			"Same weight nodes. Same chars.",
			heapOfNodes{&haffmanBTNode{weight: 10, chars: charsAB}, &haffmanBTNode{weight: 10, chars: charsAB}},
			args{0, 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("heapOfNodes.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapOfNodes_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *heapOfNodes
		want interface{}
	}{
		{
			"Pop from non-empty container.",
			&heapOfNodes{&haffmanBTNode{weight: 20}, &haffmanBTNode{weight: 10}},
			&haffmanBTNode{weight: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapOfNodes.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapOfNodes_Push(t *testing.T) {
	type args struct {
		elem interface{}
	}
	tests := []struct {
		name    string
		hBefore *heapOfNodes
		args    args
		hAfter  *heapOfNodes
	}{
		{
			"Push in non-empty container.",
			&heapOfNodes{&haffmanBTNode{weight: 20}, &haffmanBTNode{weight: 10}},
			args{&haffmanBTNode{weight: 30}},
			&heapOfNodes{&haffmanBTNode{weight: 20}, &haffmanBTNode{weight: 10}, &haffmanBTNode{weight: 30}},
		},
		{
			"Push in non-empty container.",
			&heapOfNodes{},
			args{&haffmanBTNode{weight: 30}},
			&heapOfNodes{&haffmanBTNode{weight: 30}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hBefore.Push(tt.args.elem)
			if !reflect.DeepEqual(tt.hBefore, tt.hAfter) {
				t.Errorf("heapOfNodes.Push() = %v, want %v", tt.hBefore, tt.hAfter)
			}
		})
	}
}
