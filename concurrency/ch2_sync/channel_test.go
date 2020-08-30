package ch2_sync

import (
	"fmt"
	"sync"
	"testing"
)

func channelBasic() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	value, ok := <-stringStream
	fmt.Println(value, ok)
}

func channelRange() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()
	for v := range intStream {
		fmt.Println(v)
	}
}

func multipleGoroutine() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func channelOwnerTest() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}
	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func TestChannel(t *testing.T) {
	channelBasic()
	channelRange()
	multipleGoroutine()
	channelOwnerTest()
}
