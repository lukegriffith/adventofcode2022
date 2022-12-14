package queue

const initialQueueCapacity = 32

type Queue[T any] struct {
	contents     []T
	currentIndex int
}

func (q Queue[T]) Length() int {
	return len(q.contents)
}

func (q Queue[T]) Position() int {
	return q.currentIndex
}

func (q Queue[T]) IsEmpty() bool {
	return q.currentIndex == len(q.contents)
}

func (q *Queue[T]) Push(element T) {
	q.contents = append(q.contents, element)
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		contents:     make([]T, 0, initialQueueCapacity),
		currentIndex: 0,
	}
}

func NewQueueCap[T any](initialCapacity int) Queue[T] {
	return Queue[T]{
		contents:     make([]T, 0, initialCapacity),
		currentIndex: 0,
	}
}

func (q Queue[T]) Peek() T {
	return q.contents[q.currentIndex]
}

func (q Queue[T]) PeekMore(i int) T {
	return q.contents[q.currentIndex+i]
}

func (q *Queue[T]) Next() T {
	item := q.contents[q.currentIndex]
	q.currentIndex++
	return item
}
