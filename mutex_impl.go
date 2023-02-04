package contest

import "sync"

func New() Mutex {
	return &MyMutex{
		ch: make(chan struct{}),
	}
}

type MyMutex struct {
	locked int64
	ch     chan struct{}
	mu     sync.RWMutex
}

func (m *MyMutex) Lock() {
	m.mu.Lock()
	m.locked += 1
	if m.locked == 1 {
		m.mu.Unlock()
	} else {
		m.mu.Unlock()
		<-m.ch
	}
}

func (m *MyMutex) Unlock() {
	m.mu.Lock()
	m.locked -= 1
	if m.locked == 0 {
		m.mu.Unlock()
	} else {
		m.mu.Unlock()
		m.ch <- struct{}{}
	}
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	ch := make(chan struct{})
	m.mu.Lock()
	if m.locked == 0 {
		m.locked += 1
		m.mu.Unlock()
		close(ch)
	} else {
		m.mu.Unlock()
	}
	return ch
}
