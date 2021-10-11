package ch10

import "fmt"

type ntree_node struct {
	val           interface{}
	parent        *ntree_node
	left_child    *ntree_node
	right_sibling *ntree_node
	level         int
}

type ntree struct {
	root      *ntree_node
	max_child int
	queue     []*ntree_node
}

func NewNTree(n int) *ntree {
	return &ntree{
		max_child: n,
	}
}

func (n *ntree) empty() bool {
	return n.root == nil
}
func (n *ntree) insert(val interface{}) {
	if n.empty() {
		n.root = &ntree_node{
			val:           val,
			parent:        nil,
			left_child:    nil,
			right_sibling: nil,
			level:         1,
		}
		return
	}
	parent := n.root
	for {
		left_child := parent.left_child
		i := 0
		if left_child == nil {
			parent.left_child = &ntree_node{
				val:           val,
				parent:        parent,
				left_child:    nil,
				right_sibling: nil,
				level:         parent.level + 1,
			}
			return
		}
		for ; i < n.max_child-1; i++ {
			if left_child.right_sibling == nil {
				break
			}
			left_child = left_child.right_sibling
		}
		if i == n.max_child-1 {
			if parent.right_sibling != nil {
				parent = parent.right_sibling
			} else {
				parent = parent.left_child
			}
		} else {
			left_child.right_sibling = &ntree_node{
				val:           val,
				parent:        parent,
				left_child:    nil,
				right_sibling: nil,
				level:         parent.level + 1,
			}
			return
		}
	}

}
func (n *ntree) search() {}
func (n *ntree) print() {
	root := n.root
	n.enqueue(root)
	for !n.queue_empty() {
		node := n.dequeue()
		fmt.Println(node.level, node.val)
		left := node.left_child
		for ; left != nil; left = left.right_sibling {
			n.enqueue(left)
		}
	}
}

func (n *ntree) queue_empty() bool {
	return len(n.queue) == 0
}
func (n *ntree) enqueue(node *ntree_node) {
	n.queue = append(n.queue, node)
}
func (n *ntree) dequeue() (node *ntree_node) {
	if len(n.queue) == 0 {
		return nil
	}
	node = n.queue[0]
	n.queue = n.queue[1:]
	return node
}
