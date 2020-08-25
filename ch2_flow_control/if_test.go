package ch2_flow_control

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func basicIf() {
	a := 10
	if a > 5 {
		fmt.Println("bigger than 5")
	} else {
		fmt.Println("equal or less than 5")
	}
}

func functionInsideIf() {
	filepath := "hello.txt"
	if a, err := ioutil.ReadFile(filepath); err == nil {
		fmt.Printf("%s", a)
	} else {
		fmt.Println(err)
	}
	filepath = "bye.txt"
	if a, err := ioutil.ReadFile(filepath); err == nil {
		fmt.Printf("%s", a)
	} else {
		fmt.Println(err)
	}
}

func TestArray(t *testing.T) {
	//basicIf()
	functionInsideIf()
}
