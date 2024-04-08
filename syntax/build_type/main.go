package main

//
//func main() {
//	Array()
//}

import "fmt"

func main() {
	mySlice := make([]int, 0, 2) // Creating a slice with length 0 and capacity 5

	//fmt.Printf("Type of myInt: %v\n", reflect.TypeOf(mySlice))
	//fmt.Printf("Type of myInt: %v\n", reflect.TypeOf(&mySlice))
	fmt.Printf("Original slice address: %p, Capacity: %d, Length: %d\n", mySlice, cap(mySlice), len(mySlice))

	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("After appending %d, slice address: %p, Capacity: %d, Length: %d\n", i, mySlice, cap(mySlice), len(mySlice))
	}
}
