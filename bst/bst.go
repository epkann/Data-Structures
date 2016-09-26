package bst

// bst implements a non-self-balancing binary search tree.
// All values less than that of a node are in the left sub-tree;
// values greater than or equal are in the right sub-tree.

// Value must be comparable by a less method.
type Value interface {
	Less(other Value) bool
}

// Node is a single element within the tree.
type Node struct {
	Value
	parent *Node
	left   *Node
	right  *Node
}

// Tree holds elements of binary search tree.
type Tree struct {
	head *Node
	size int
}

// Size returns number of nodes in the tree.
func (tree *Tree) Size() int {
	return tree.size
}

// Search returns true if value is in the tree, false otherwise.
func (tree *Tree) Search(v Value) bool {
	h := tree.head

	// is there a way to make the following a switch statement?
	for h != nil {
		if !h.Value.Less(v) && !v.Less(h.Value) {
			return true
		} else if h.Value.Less(v) {
			h = h.right
		} else {
			h = h.left
		}
	}
	return false
}

// Insert inserts a single element in the tree.
func (tree *Tree) Insert(v Value) {
	n := &Node{Value: v}
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
			} else {
				h = h.left
			}
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
	return
}

// Delete deletes the node with the given value.
// If multiple nodes contain the same value, one is deleted arbitrarily.
func (tree *Tree) Delete(v Value) {

	h := tree.head
	for h != nil {
		if !v.Less(h.Value) && !h.Value.Less(v) { // found node to delete
			if h.left == nil {
				tree.Transplant(h, h.right)
			} else if h.right == nil {
				tree.Transplant(h, h.left)
			} else {
				y := Minimum(h.right) // find successor node
				if y.parent != h {
					tree.Transplant(y, y.right)
					y.right = h.right
					y.right.parent = y
				}
				tree.Transplant(h, y)
				y.left = h.left
				y.left.parent = y
			}
			tree.size--
			return
		} else if v.Less(h.Value) {
			h = h.left
		} else {
			h = h.right
		}
	}
	return
}

// Transplant replaces the subtree rooted at node u with the subtree rooted at node v.
// Does not update v.left or v.right.
func (tree *Tree) Transplant(u, v *Node) {
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
func Minimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// InOrder returns all values in-order.
func (tree *Tree) InOrder() []Value {
	values := make([]Value, 0, tree.size)
	n := tree.head
	Traverse(n, &values)
	return values
}

// Traverse traverses a tree in-order starting from the passed node.
func Traverse(n *Node, values *[]Value) {
	if n == nil {
		return
	}
	Traverse(n.left, values)
	*values = append(*values, n.Value)
	Traverse(n.right, values)
	return
}
