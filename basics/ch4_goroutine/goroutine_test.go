package ch4_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func basicGoroutine() {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	fmt.Println("end loop")
}

func sum(a, b int, c chan int) {
	c <- a + b
}

func basicChannel() {
	c := make(chan int)
	go sum(1, 2, c)
	n := <-c
	fmt.Println(n)
}

// 체널의 버퍼 크기를 설정하지 않고 그냥 생성하면 동기 체널
// 동기 체널은 보내는 쪽에서 값을 받을 때까지 대기, 받는 쪽에서는 체널에 값이 들어올 때까지 대기
// 동기 체널 활용하면 고루틴의 코드 실행 순서 제어 가능
func syncChannel() {
	count := 1000
	done := make(chan bool)
	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("goroutine ", i)
		}
	}()
	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main: ", i)
	}
}

func asyncChannel() {
	count := 1000
	done := make(chan bool, 5)
	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("goroutine ", i)
		}
	}()
	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main: ", i)
	}
}

// range라는 구문은 체널이 닫힐 때까지 값을 꺼낸다.
// 그러므로 체널을 닫지 않은 채로 range문을 호출하게 되면 마지막 단계에서 deadlock에 걸리게 된다.
func syncRange() {
	c := make(chan int)
	count := 1000
	go func() {
		for i := 0; i < count; i++ {
			c <- i
			fmt.Println("goroutine: ", i)
		}
		close(c)
	}()
	for i := range c {
		fmt.Println("main: ", i)
	}
}

func asyncRange() {
	c := make(chan int, 10)
	count := 1000
	go func() {
		for i := 0; i < count; i++ {
			c <- i
			fmt.Println("goroutine: ", i)
		}
		close(c)
	}()
	for i := range c {
		fmt.Println("main: ", i)
	}
}

// chan<-: send only channel
func sendOnlyParam(c chan<- int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	// send only 체널에서 값을 꺼내면 compile error 발생
	// fmt.Println(<-c)
}

func receiveOnlyParam(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
	// receive only 체널에 값을 넣으면 compile error 발생
	// c<-100
}

func receiveOnlyReturn() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func sumReceiveOnly(c <-chan int) <-chan int {
	out := make(chan int)
	r := 0
	go func() {
		for i := range c {
			r = r + i
			out <- r
		}
		close(out)
	}()
	return out
}

func receiveOnlyTest() {
	c1 := receiveOnlyReturn()
	c2 := sumReceiveOnly(c1)
	for i := range c2 {
		fmt.Println(i)
	}
}

func TestStruct(t *testing.T) {
	runtime.GOMAXPROCS(1)
	//basicGoroutine()
	//basicChannel()
	//syncChannel()
	//asyncChannel()
	syncRange()
	//asyncRange()
	//sendOnlyReturnTest()
	//receiveOnlyTest()
}
