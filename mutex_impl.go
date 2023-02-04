package contest

import "sync/atomic"

func New() Mutex {
	return &MyMutex{
		ch: make(chan struct{}),
	}
}

type MyMutex struct {
	locked int32
	count  int64
	ch     chan struct{}
}

func (m *MyMutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.locked, 0, 1) {
		atomic.AddInt64(&m.count, 1)
	} else {
		atomic.AddInt64(&m.count, 1)
		<-m.ch
	}
}

func (m *MyMutex) Unlock() {
	if atomic.CompareAndSwapInt64(&m.count, 1, 0) {
		atomic.AddInt32(&m.locked, -1)
	} else {
		atomic.AddInt64(&m.count, -1)
		m.ch <- struct{}{}
	}
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{})
	if atomic.CompareAndSwapInt32(&m.locked, 0, 1) {
		atomic.AddInt64(&m.count, 1)
		close(ch)
	}
	return ch
}
