package util

import "container/heap"

type PriorityQueue[T any] struct {
	inner *heapInterface[T]
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		inner: &heapInterface[T]{
			less: less,
		},
	}
}

func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(pq.inner, item)
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(pq.inner).(T)
}

type heapInterface[T any] struct {
	arr  []T
	less func(a, b T) bool
}

var _ heap.Interface = &heapInterface[struct{}]{}

func (pq *heapInterface[T]) Push(x any) {
	pq.arr = append(pq.arr, x.(T))
}

func (pq *heapInterface[T]) Pop() any {
	lastIdx := len(pq.arr) - 1

	item := pq.arr[lastIdx]
	pq.arr = pq.arr[:lastIdx]

	return item
}

func (pq *heapInterface[T]) Len() int {
	return len(pq.arr)
}

func (pq *heapInterface[T]) Less(i, j int) bool {
	a, b := pq.arr[i], pq.arr[j]

	return pq.less(a, b)
}

func (pq *heapInterface[T]) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}
