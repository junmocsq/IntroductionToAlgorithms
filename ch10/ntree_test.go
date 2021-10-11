package ch10

import "testing"

func TestNewNTree(t *testing.T) {
	tree := NewNTree(4)
	for i:=100;i<130;i++{
		tree.insert(i)
	}
	tree.print()
}
