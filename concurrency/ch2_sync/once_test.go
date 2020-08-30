package ch2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func onceBasic() {
	var cnt int
	var once sync.Once

	increment := func() {
		cnt++
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Println("once result: ", cnt)
}

func onceDo() {
	var cnt int
	increment := func() { cnt++ }
	decrement := func() { cnt-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", cnt)
}

func TestOnce(t *testing.T) {
	onceBasic()
	onceDo()
}
