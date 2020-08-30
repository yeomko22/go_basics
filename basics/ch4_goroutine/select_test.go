package ch4_goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func multipleChannelSelect() {
	c1 := make(chan int)
	c2 := make(chan int)
	cnt := 1000
	go func() {
		for i := 0; i < cnt; i++ {
			c1 <- i
		}
	}()

	go func() {
		for i := 0; i < cnt; i++ {
			c2 <- i
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println("ch1: ", i)
			case i := <-c2:
				fmt.Println("ch2: ", i)
			case <-time.After(50 * time.Millisecond):
				fmt.Println("timeout")
			}
		}
	}()
	time.Sleep(10 * time.Second)
}

func sendInSelect() {
	c1 := make(chan int)
	go func() {
		for {
			select {
			case c1 <- 10:
			case i := <-c1:
				fmt.Println("c1", i)
			}
		}
	}()
	time.Sleep(1 * time.Second)
}

func TestSelect(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//multipleChannelSelect()
	sendInSelect()
}
