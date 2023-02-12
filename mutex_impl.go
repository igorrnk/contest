package contest

func New() Mutex {
	m := new(MyMutex)
	m.ch = make(chan struct{})
	sync := make(chan bool)
	go m.writeToChannel(sync)
	<-sync
	return m
}

type MyMutex struct {
	ch chan struct{}
}

func (m *MyMutex) Lock() {
	<-m.ch
}

func (m *MyMutex) Unlock() {
	sync := make(chan bool)
	go m.writeToChannel(sync)
	<-sync
}

func (m *MyMutex) LockChannel() <-chan struct{} {
	return m.ch
}

func (m *MyMutex) writeToChannel(sync chan bool) {
	close(sync)
	m.ch <- struct{}{}
}
