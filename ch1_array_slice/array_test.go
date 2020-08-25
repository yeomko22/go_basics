package ch1_array_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func createArray() {
	var a [5]int
	a[2] = 1
	fmt.Println("a", a, reflect.TypeOf(a))

	// int 배열 선언과 동시에 값 할당
	b := [5]int{1, 2, 3}
	fmt.Println("b: ", b, reflect.TypeOf(b))

	// ...을 사용하여 배열의 크기를 초기화 값의 개수에 맞춰서 동적으로 결정
	c := [...]int{1, 2, 3}
	fmt.Println("c: ", c, reflect.TypeOf(c))
}

func compareEmptyArrayType() {
	a := [1]int{}
	b := [3]int{}
	fmt.Println(a, reflect.TypeOf(a))
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
}

func scanArray() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, e := range a {
		fmt.Println(i, e)
	}
}

func copyArray() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a
	b[1] = 6
	fmt.Println(a)
	fmt.Println(b)
}

func TestArray(t *testing.T) {
	createArray()
	//compareEmptyArrayType()
	//scanArray()
	//copyArray()
}
