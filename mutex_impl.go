package contest

import "sync"

func New() Mutex {
	return &MutexImpl{}
}

type MutexImpl struct {
	sync.Mutex
}

func (m *MutexImpl) LockChannel() <-chan struct{} {
	c := make(chan struct{})
	if m.TryLock() {
		close(c)
	}
	return c
}
