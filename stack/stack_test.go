package stack

import "testing"

func TestIntStack(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	expected := []int{1, 2, 3}
	if got := s.stack; got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Errorf("Got %v expected %v", got, expected)
	}
	s.Pop()
	expected = []int{1, 2}
	if got := s.stack; got[0] != 1 || got [1] != 2 || len(got) != 2 {
			t.Errorf("Got %v expected %v", got, expected)
	}
}