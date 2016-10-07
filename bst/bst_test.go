package bst

import "testing"

type Int int

func (i Int) Less(other Value) bool {
	return i < other.(Int)
}

func TestTree(t *testing.T) {
	tree := new(Tree)
	tree.Insert(Int(8))
	tree.Insert(Int(5))
	tree.Insert(Int(2))
	tree.Insert(Int(3))
	tree.Insert(Int(11))
	tree.Insert(Int(13))
	tree.Insert(Int(1))
	tree.Insert(Int(12))
	if actualSize := tree.Len(); actualSize != 8 {
		t.Errorf("Got %v expected 8", actualSize)
	}
	if actual := tree.Search(Int(5)); actual != true {
		t.Errorf("Got %v expected true", actual)
	}
	if actualBool := tree.Search(Int(16)); actualBool != false {
		t.Errorf("Got %v expected false", actualBool)
	}
	tree.Delete(Int(8))
	if actualAfterDelete := tree.Search(Int(8)); actualAfterDelete != false {
		t.Errorf("Got %v expected false", actualAfterDelete)
	}
	if sizeAfterDelete := tree.Len(); sizeAfterDelete != 7 {
		t.Errorf("Got %v expected 7", sizeAfterDelete)
	}
}
