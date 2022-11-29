package queueforstack

type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(ele T) {
	s.elements = append(s.elements, ele)
}

func (s *stack[T]) pop() T {
	var ele T
	if s.empty() {
		return ele
	}
	ele = s.elements[s.len()-1]
	s.elements = s.elements[:s.len()-1]
	return ele
}

func (s *stack[T]) len() int {
	return len(s.elements)
}

func (s *stack[T]) empty() bool {
	return s.len() == 0
}

type stackQueue[T any] struct {
	in  stack[T]
	out stack[T]
}

func (q *stackQueue[T]) push(t T) {
	q.in.push(t)
}

func (q *stackQueue[T]) pop() (t T) {
	if !q.out.empty() {
		return q.out.pop()
	}
	for !q.in.empty() {
		q.out.push(q.in.pop())
	}
	if q.out.empty() {
		return t
	}
	return q.out.pop()
}
