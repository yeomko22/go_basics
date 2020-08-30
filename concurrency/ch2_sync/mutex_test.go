package ch2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func WithMutexTest() int {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 20; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 20; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	return count
}

func WithoutMutexTest() int {
	var count int

	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 20; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 20; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	return count
}

func WithWithoutMutexTest() {
	result := 0
	for i := 0; i < 1000; i++ {
		testResult := WithMutexTest()
		if testResult != 0 {
			result++
		}
	}
	fmt.Printf("with mutex, trial: %d, error: %d\n", 1000, result)

	result = 0
	for i := 0; i < 1000; i++ {
		testResult := WithoutMutexTest()
		if testResult != 0 {
			result++
		}
	}
	fmt.Printf("without mutex trial: %d, error: %d\n", 1000, result)
}

func TestMutex(t *testing.T) {
	WithWithoutMutexTest()
}
