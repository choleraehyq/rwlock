package rwlock

import (
	"sync"

	"github.com/jonhoo/drwmutex"
)

type RWLock struct {
	drwmutex.DRWMutex
}

func New() *RWLock {
	return &RWLock{
		drwmutex.New(),
	}
}

func (this *RWLock) Lock() {
	this.DRWMutex.Lock()
}

func (this *RWLock) Unlock() {
	this.DRWMutex.Unlock()
}

func (this *RWLock) RLocker() sync.Locker {
	return this.DRWMutex.RLocker()
}

