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

func init() {
	shardsLen = runtime.GOMAXPROCS(0)
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

func (this RWLock) RLock() {
	return this.RLocker().Lock()
}

func (this RWLock) RUnlock() {
	return this.RLocker().Unlock()
}
