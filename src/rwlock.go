package main

import "sync"

//RWLock some more
type RWLock struct {
	// synchronziation variables
	mu      *sync.Mutex
	readGo  *sync.Cond
	writeGo *sync.Cond

	// state variables
	activeReaders, waitingReaders int
	activeWriters, waitingWriters int
}

// NewRWLock returns an RWLock
func NewRWLock() *RWLock {
	rwl := &RWLock{}
	rwl.mu = &sync.Mutex{}
	rwl.readGo = sync.NewCond(rwl.mu)
	rwl.writeGo = sync.NewCond(rwl.mu)
	return rwl
}

func (rw *RWLock) startRead() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.waitingReaders++
	for rw.readShouldWait() {
		rw.readGo.Wait()
	}
	rw.waitingReaders--
	rw.activeReaders++
}

func (rw *RWLock) doneRead() {

}

func (rw *RWLock) startWrite() {

}

func (rw *RWLock) doneWrite() {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.activeWriters--
	// assert activeWriters == 0
	if rw.waitingWriters > 0 {
		rw.writeGo.Signal()
	} else {
		// here there may be waiting readers
		rw.readGo.Broadcast()
	}
}

func (rw *RWLock) readShouldWait() bool {

}

func (rw *RWLock) writeShouldWait() bool {
}
