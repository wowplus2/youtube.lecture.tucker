package dataStruct

type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {
	return &Queue{ll: &LinkedList{}}
}

func (q *Queue) Push(n int) {
	q.ll.AddNode(n)
}

func (q *Queue) Pop() int {
	front := q.ll.Front()
	q.ll.PopFront()

	return front
}

func (q *Queue) IsEmpty() bool {
	return q.ll.IsEmpty()
}
