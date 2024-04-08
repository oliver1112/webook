package main

import "fmt"

func Array() {
	a1 := [5]int{9, 8, 7}
	fmt.Printf("a1: %v, len = %d, cap = %d ", a1, len(a1), cap(a1))
}
