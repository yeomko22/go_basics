package ch2_flow_control

import (
	"fmt"
	"testing"
)

func basicFor() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func nestedForWithLabel() {
Loop:
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < 5; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println("j: ", j)
		}
	}
}

func basicGoto() {
	a := 1
	if a == 1 {
		goto ERROR1
	} else {
		goto ERROR2
	}
ERROR1:
	fmt.Println("error 1")
	return
ERROR2:
	fmt.Println("error 2")
	return
}

func TestFor(t *testing.T) {
	//basicFor()
	//nestedForWithLabel()
	basicGoto()
}
