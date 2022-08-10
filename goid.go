//go:build amd64 || arm64 || arm || 386 || mipsle
// +build amd64 arm64 arm 386 mipsle

package rwlock

func goid() int
