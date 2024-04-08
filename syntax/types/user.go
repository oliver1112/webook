package main

import "fmt"

func NewUser() {
	u := User{}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)

	up := &User{}
	fmt.Printf("%+v\n", up)

	//up3 := &User{Name: "Jennifer", Age: 20}
	//println(up3.Name)
}

type User struct {
	Name string
	Age  int
}
