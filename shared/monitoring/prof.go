package monitoring

// https://golang.org/pkg/runtime/pprof/#Profiles

import (
	"runtime/pprof"
)

func getCount(name string) int {
	p := pprof.Lookup(name)
	if p == nil {
		return 0
	}
	return p.Count()
}

func countGoroutine() int {
	return getCount("goroutine")
}

func countThreadCreate() int {
	return getCount("threadcreate")
}

func countHeap() int {
	return getCount("heap")
}

func countBlock() int {
	return getCount("block")
}

func countMutex() int {
	return getCount("mutex")
}
