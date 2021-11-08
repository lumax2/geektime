package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(0)
	for i := 0; i < b.N; i++ {
		vv := v.Load().(int)
		vv++
		v.Store(vv)
	}
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	cnt := 0
	for i := 0; i < b.N; i++ {
		mu.Lock()
		cnt++
		mu.Unlock()
	}
}

func BenchmarkRWMutex(b *testing.B) {
	var mu sync.RWMutex
	cnt := 0
	for i := 0; i < b.N; i++ {
		mu.Lock()
		cnt++
		mu.Unlock()
	}
}
