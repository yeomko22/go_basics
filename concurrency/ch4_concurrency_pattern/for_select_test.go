package ch4_concurrency_pattern

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func forSelectBasic() {
	done := make(chan int)
	for {
		select {
		case <-done:
		default:
			{
				fmt.Println("default")
				time.Sleep(1 * time.Second)
			}
			fmt.Println("for loop")
			time.Sleep(1 * time.Second)
		}
	}
}

//func preventMemoryLeak() {
//	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
//		terminated := make(chan interface{})
//		go func() {
//			defer fmt.Println("doWork exited.")
//			defer close(terminated)
//			for {
//				select {
//				case s := <-strings:
//					fmt.Println(s)
//				case <-done:
//					fmt.Println("receive done signal")
//					return
//				default:
//					fmt.Println("default")
//					time.Sleep(1 * time.Second)
//				}
//				fmt.Println("for loop")
//				time.Sleep(1 * time.Second)
//			}
//		}()
//		return terminated
//	}
//	done := make(chan interface{})
//	terminated := doWork(done, nil)
//	go func() {
//		time.Sleep(5 * time.Second)
//		fmt.Println("Canceling doWork goroutine...")
//		close(done)
//	}()
//	<-terminated
//	fmt.Println("Done.")
//}

func preventWriteToChannel() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}

			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func TestForSelect(t *testing.T) {
	//preventMemoryLeak()
	preventWriteToChannel()
}
