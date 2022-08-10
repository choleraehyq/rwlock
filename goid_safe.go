//go:build !amd64 && !arm64 && !arm && !386 && !mipsle
// +build !amd64,!arm64,!arm,!386,!mipsle

package rwlock

import (
	"runtime"
)

func goid() (n int) {
	const offset = len("goroutine ")
	var data [32]byte
	b := data[:runtime.Stack(data[:], false)]
	if len(b) <= offset {
		return
	}
	for i := offset; i < len(b); i++ {
		j := int(b[i] - '0')
		if j < 0 || j > 9 {
			break
		}
		n = n*10 + j
	}
	return n
}
