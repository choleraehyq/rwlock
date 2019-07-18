package rwlock

import (
	pid "github.com/choleraehyq/pid"
	"runtime"
	"sync"
)

const (
	cacheLineSize = 64
)

var (
	shardsLen int
)

type RWLock []rwmutexShard

type rwmutexShard struct {
	_ [cacheLineSize]byte
	sync.RWMutex
}

func New() RWLock {
	shardsLen = runtime.GOMAXPROCS(0)
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

