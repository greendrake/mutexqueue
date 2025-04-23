package mutexqueue

import (
	"sync"
)

type MutexQueue struct {
	items []interface{}
	lock  sync.Mutex
}

func New() *MutexQueue {
	return &MutexQueue{
		items: []interface{}{},
	}
}

func (q *MutexQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *MutexQueue) GetSize() int {
	return len(q.items)
}

func (q *MutexQueue) Put(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

func (q *MutexQueue) Get() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
