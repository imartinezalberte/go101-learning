// race conditions
// A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this
// order is guaranteed to be maintained.
package chapter1

import (
	"sync"
	"time"
)

const (
	DataIsZero    = "data is zero"
	DataIsNotZero = "data is not zero"
)

func Example1() string {
	var data int
	go func() {
		data++ // three operation here, but for the shake of simplicity, it's just one
	}()
	// Another operation that should execute after the goroutine takes place, but the likelihood that is gonna happen, is way lesser than expected.
	if data == 0 {
		return DataIsZero // not really, because the data value can be changed while executing the if statement. Only god knows...
	}
	return DataIsNotZero
}

func Example2ErroneousSolution() string {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second)
	if data == 0 {
		return DataIsZero
	}
	return DataIsNotZero
}

func Example3WaitGroupSolution() string {
	var (
		data int
		wg *sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()
	if data == 0 {
		return DataIsZero
	}
	return DataIsNotZero
}

