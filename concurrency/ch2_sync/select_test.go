package ch2_sync

import (
	"fmt"
	"testing"
	"time"
)

func selectMultiple() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1Count: %d\n c2Count: %d\n", c1Count, c2Count)
}

func selectFor() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

func TestSelect(t *testing.T) {
	//selectMultiple()
	selectFor()
}
