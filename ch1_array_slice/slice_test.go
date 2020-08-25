package ch1_array_slice

import (
	"fmt"
	"reflect"
	"testing"
)

func createSlice() {
	var a []int
	fmt.Println("a", a, reflect.TypeOf(a), a == nil)

	b := []int{}
	fmt.Println(b, reflect.TypeOf(b), b == nil)
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))

	c := make([]int, 5)
	fmt.Println("c", b, reflect.TypeOf(c))

	d := []int{1, 2, 3, 4, 5}
	fmt.Println("d", c, reflect.TypeOf(d))

	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(b) == reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(c) == reflect.TypeOf(d))
}

func addValueToSlice() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6)
	fmt.Println(a)

	b := []int{7, 8, 9}
	a = append(a, b...)
	fmt.Println(a)
}

func copySlice() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 5
	fmt.Println("a", a)
	fmt.Println("b", b)

	c := []int{1, 2, 3}
	d := make([]int, len(c))
	copy(d, c)
	d[0] = 5
	fmt.Println("c", c)
	fmt.Println("d", d)
}

func sliceCapacity() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(len(a), cap(a))
	a = append(a, 6, 7)
	fmt.Println(len(a), cap(a))
}

func subSlice() {
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
	b[0] = 7
	fmt.Println(a)
	fmt.Println(b)
}

func TestSlice(t *testing.T) {
	createSlice()
	addValueToSlice()
	copySlice()
	sliceCapacity()
	subSlice()
}
