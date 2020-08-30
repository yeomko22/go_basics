package ch2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func waitgroupBasic() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func TestWaitGroup(t *testing.T) {
	waitgroupBasic()
}
