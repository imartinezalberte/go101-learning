// questions to answer:
// Are my critical sections entered and exited repeatedly?
// What size should my critical sections be?
//
// A deadlocked program is one in which all concurrent processes are waiting on one another.
//
// 1971, Edgar Coffman enumared these conditions:
//
// - Mutual Exclusion: A concurrent process holds exclusive rights to a resource at any one time.
// - Wait For Condition: A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
// - No Preemption: A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
// - Circular Wait: A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn waiting on it (P1), so it fulfills this final condition too.
package chapter1

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	sync.Mutex
	value int
}

func Deadlock() {
	var (
		wg sync.WaitGroup
		a, b value
		f = func(a *value, b *value) {
			defer wg.Done()
			a.Lock()
			defer a.Unlock()

			time.Sleep(2 * time.Second) // We shouldn't do that way, but it's useful right now...

			b.Lock()
			defer b.Unlock()
			fmt.Println(a.value + b.value)
		}
	)
	wg.Add(2)	
	defer wg.Wait()
	go f(&a, &b)
	go f(&b, &a)
}
