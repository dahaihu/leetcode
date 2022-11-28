package queueforstack

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Push(ele T) {
	q.elements = append(q.elements, ele)
}

func (q *Queue[T]) Pop() (t T) {
	if q.Empty() {
		return
	}
	t = q.elements[0]
	q.elements = q.elements[1:]
	return
}

func (q *Queue[T]) Len() int {
	return len(q.elements)
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

type Stack[T any] struct {
	use  Queue[T]
	help Queue[T]
}

func (s *Stack[T]) Push(t T) {
	if s.use.Empty() {
		s.use.Push(t)
		return
	}
	s.help.Push(t)
	for !s.use.Empty() {
		s.help.Push(s.use.Pop())
	}
	s.use, s.help = s.help, s.use
}

func (s *Stack[T]) Pop() T {
	return s.use.Pop()
}

func (s *Stack[T]) Empty() bool {
	return s.use.Empty()
}
