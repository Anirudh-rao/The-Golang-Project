package Extras

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounterFunction() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for range 50 {
		wg.Go(func() {
			ops.Add(1)
		})

	}
	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
