// Package bst implements a non-self-balancing binary search tree.
package bst

// Value must be comparable by a less method.
type Value interface {
	Less(other Value) bool
}

// Node is a single element within the tree.
type node struct {
	Value
	parent *node
	left   *node
	right  *node
}

// Tree holds elements of non-self-balancing binary search tree.
// All values less than that of a node are in the left sub-tree;
// values greater than or equal are in the right subtree.
type Tree struct {
	head *node
	size int
}

// Len returns number of nodes in the tree.
func (tree *Tree) Len() int {
	return tree.size
}

// Search returns whether the value is in the tree.
func (tree *Tree) Search(v Value) bool {
	h := tree.head

	// TODO(epkann): is there a way to make the following a switch statement?
	for h != nil {
		if !h.Value.Less(v) && !v.Less(h.Value) {
			return true
		} 
		if h.Value.Less(v) {
			h = h.right
		} else {
			h = h.left
		}
	}
	return false
}

// Insert inserts a single element in the tree.
func (tree *Tree) Insert(v Value) {
	n := &node{Value: v}
	if tree.head == nil {
		tree.head = n
		tree.size++
		return
	}

	h := tree.head

	for {
		if n.Value.Less(h.Value) {
			if h.left == nil {
				h.left = n
				n.parent = h
				break
			} 
			h = h.left
		} else {
			if h.right == nil {
				h.right = n
				n.parent = h
				break
			} else {
				h = h.right
			}
		}
	}
	tree.size++
}

// Delete deletes the node with the given value.
// If multiple nodes contain the same value, one is deleted arbitrarily.
func (tree *Tree) Delete(v Value) {
	h := tree.head
	for h != nil {
		if !v.Less(h.Value) && !h.Value.Less(v) { // found node to delete
			switch {
			case h.left == nil:
				tree.transplant(h, h.right)
			case h.right == nil:
				tree.transplant(h, h.left)
			default: 
				y := minimum(h.right) // find successor node
				if y.parent != h {
					tree.transplant(y, y.right)
					y.right = h.right
					y.right.parent = y
				}
				tree.transplant(h, y)
				y.left = h.left
				y.left.parent = y
			}
			tree.size--
			return
		} 
		if v.Less(h.Value) {
			h = h.left
		} else {
			h = h.right
		}
	}
}

// Transplant replaces the subtree rooted at node u with the subtree rooted at node v.
// Does not update v.left or v.right.
func (tree *Tree) transplant(u, v *node) {
	if u.parent == nil {
		tree.head = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Minimum returns the minimum valued node of the subtree rooted at n.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// InOrder returns all values in-order.
func (tree *Tree) InOrder() []Value {
	values := make([]Value, 0, tree.size)
	n := tree.head
	traverse(n, &values)
	return values
}

// Traverse traverses a tree in-order starting from the passed node.
func traverse(n *node, values *[]Value) {
	if n == nil {
		return
	}
	traverse(n.left, values)
	*values = append(*values, n.Value)
	traverse(n.right, values)
	return
}
