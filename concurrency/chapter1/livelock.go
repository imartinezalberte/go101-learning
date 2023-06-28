// Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move
// the state of the program forward.
package chapter1

import (
	"sync"
	"time"
)

var cadence = sync.NewCond(new(sync.Mutex))

func E() {
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()
}
