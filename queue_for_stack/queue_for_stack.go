package queueforstack

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) push(ele T) {
	q.elements = append(q.elements, ele)
}

func (q *Queue[T]) pop() (t T) {
	if q.empty() {
		return
	}
	t = q.elements[0]
	q.elements = q.elements[1:]
	return
}

func (q *Queue[T]) len() int {
	return len(q.elements)
}

func (q *Queue[T]) empty() bool {
	return q.len() == 0
}

type queueStack[T any] struct {
	use  Queue[T]
	help Queue[T]
}

func (s *queueStack[T]) push(t T) {
	if s.use.empty() {
		s.use.push(t)
		return
	}
	s.help.push(t)
	for !s.use.empty() {
		s.help.push(s.use.pop())
	}
	s.use, s.help = s.help, s.use
}

func (s *queueStack[T]) pop() T {
	return s.use.pop()
}

func (s *queueStack[T]) empty() bool {
	return s.use.empty()
}
