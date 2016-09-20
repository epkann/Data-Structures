package stack

type Stack struct {
	stack []interface{}
	len   int
}

func (s *Stack) Pop() (item interface{}) {
	item = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.len--
	return
}

func (s *Stack) Push(item interface{}) {
	s.stack = append(s.stack, item)
	s.len++
	return
}

func (s *Stack) Peek() (item interface{}) {
	return s.stack[len(s.stack)-1]
}


func (s *Stack) Len() int {
	return s.len
}
