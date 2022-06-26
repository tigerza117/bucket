package bucket

import (
	"sync"
)

type Bucket[T any] interface {
	Add(T)
	List() []T
}

type bucket[T any] struct {
	data    []T
	size    int
	maxSize int
	sync    sync.Locker
}

func New[T any](maxSize int) Bucket[T] {
	bk := bucket[T]{
		data:    make([]T, maxSize),
		size:    0,
		maxSize: maxSize,
		sync:    NewSpinLock(),
	}

	return &bk
}

func (b bucket[T]) Add(data T) {
	b.sync.Lock()
	if b.size < b.maxSize {
		b.data[b.size] = data
		b.size++
	} else {
		b.data = b.data[1:]
		b.data = append(b.data, data)
	}
	b.sync.Unlock()
}

func (b bucket[T]) List() []T {
	return b.data
}
