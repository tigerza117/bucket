package bucket_list

import (
	"sync"
)

type Bucket[T any] interface {
	Add(T)
	List()
}

type bucket[T any] struct {
	data    []T
	size    int
	maxSize int
	sync    sync.Locker
}

func (b *bucket[T]) Add(num T) {
	b.sync.Lock()
	if b.size < b.maxSize {
		b.data[b.size] = num
		b.size++
	} else {
		b.data = b.data[1:]
		b.data = append(b.data, num)
	}
	b.sync.Unlock()
}

func (b *bucket[T]) List() []T {
	return b.data
}
