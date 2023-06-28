package chapter1

import (
	"fmt"
	"sync"
)

func ExampleCriticalSections() {
	var data int
	go func() {
		data++ // critical section because we are accessing/modifying the shared resource.
	}()
	if data == 0 { // another critical section.
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %d\n", data) // another critical section.
	}
}

// we are not waiting until the goroutine executes, we are only handling the access to the data variable
// So the function keeps being nondeterministic
func ExampleWithMutex() {
	var (
		memoryAccess sync.Mutex
		data         int
	)
	go func() { // wrapping the critical section
		memoryAccess.Lock()
		defer memoryAccess.Unlock()
		data++
	}()
	memoryAccess.Lock()
	defer memoryAccess.Unlock()
	if data == 0 { // wrapping critical section
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %d\n", data) // wrapping critical section
	}
}
