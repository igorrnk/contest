package contest

import "sync"

func New() Mutex {
	return &MutexImpl{}
}

type MutexImpl struct {
	sync.Mutex
}

func (m *MutexImpl) Lock() {
	m.Mutex.Lock()
}

func (m *MutexImpl) Unlock() {
	m.Mutex.Unlock()
}

func (m *MutexImpl) LockChannel() <-chan struct{} {
	c := make(chan struct{})
	if m.TryLock() {
		close(c)
	}
	return c
}
