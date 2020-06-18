package SortedBinaryTree

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestBinaryTree_InsertItemAndSearchItem(t *testing.T) {
	tree := SortedBinaryTree{
		root: &Node{
			data:  rand.Int(),
			left:  nil,
			right: nil,
		},
	}

	var a []int
	var b int
	for i := 0; i <= 1000; i++ {
		b = rand.Int()
		tree.InsertItem(b)
		a = append(a, b)
	}

	for _, v := range a {
		if _, in := tree.SearchItem(v); !in {
			t.Fatal()
		}
	}
}

func TestSortedBinaryTree_GetRoot(t *testing.T) {
	node := Node{
		data: rand.Int(),
	}
	tree := SortedBinaryTree{
		root: &node,
	}

	if tree.GetRoot().data != node.data {
		t.Fatal(reflect.DeepEqual(tree.GetRoot(), node))
	}
}


