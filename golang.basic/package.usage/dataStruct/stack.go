package dataStruct

type Stack struct {
	ll *LinkedList
}

func (s *Stack) IsEmpty() bool {
	return s.ll.IsEmpty()
}

func NewStack() *Stack {
	return &Stack{ll: &LinkedList{}}
}

func (s *Stack) Push(n int) {
	s.ll.AddNode(n)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()

	return back
}
