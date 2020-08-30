package ch2_flow_control

import (
	"fmt"
	"testing"
)

func basicSwitch(num int) {
	switch num {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
		fallthrough
	case 3:
		fmt.Println("TRHEE")
	}
}

func multipleCases(num int) {
	switch num {
	case 1, 2:
		fmt.Println("ONE or TWO")
	case 3:
		fmt.Println("TRHEE")
	}
}

func conditionInsideCase(num int) {
	switch {
	case num < 5:
		fmt.Println("UNDER FIVE")
	case num == 5:
		fmt.Println("EQUAL FIVE")
	case num > 5:
		fmt.Println("OVER FIVE")
	}
}

func TestSwitch(t *testing.T) {
	basicSwitch(1)
	basicSwitch(2)
	multipleCases(1)
	multipleCases(2)
	conditionInsideCase(4)
	conditionInsideCase(5)
	conditionInsideCase(6)
}
