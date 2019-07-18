package rwlock

import (
	"sync"
)

type RWLock struct {
	sync.RWMutex
}

func New() *RWLock {
	return &RWLock{}
}

func (this *RWLock) Lock() {
	this.RWMutex.Lock()
}

func (this *RWLock) Unlock() {
	this.RWMutex.Unlock()
}

func (this *RWLock) RLocker() sync.Locker {
	return this.RWMutex.RLocker()
}

