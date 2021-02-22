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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("heapOfNodes.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapOfNodes_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    heapOfNodes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_heapOfNodes_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *heapOfNodes
		want interface{}
	}{
		// TODO: Add test cases.
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
		name string
		h    *heapOfNodes
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.elem)
		})
	}
}
