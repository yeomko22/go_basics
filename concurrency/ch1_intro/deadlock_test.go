package ch1_intro

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type value struct {
	mu  sync.Mutex
	val int
}

func deadlockTest() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(1 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Println("sum: ", v1.val+v2.val)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func TestDeadlock(t *testing.T) {
	deadlockTest()
}
