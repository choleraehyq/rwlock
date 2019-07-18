package rwlock

import (
	pid "github.com/choleraehyq/pid"
	"sync"
)

const (
	cacheLineSize = 64
	shardsLen = 64
)

type RWLock []rwmutexShard

type rwmutexShard struct {
	_ [cacheLineSize]byte
	sync.RWMutex
}

func New() RWLock {
	return RWLock(make([]rwmutexShard, shardsLen))
}

func (this RWLock) Lock() {
	for shard := range this {
		this[shard].Lock()
	}
}

func (this RWLock) Unlock() {
	for shard := range this {
		this[shard].Unlock()
	}
}

func (this RWLock) RLocker() sync.Locker {
	tid := pid.GetPid()
	return this[tid % shardsLen].RWMutex.RLocker()
}

