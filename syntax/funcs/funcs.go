package main

import (
	"fmt"
	"strconv"
)

func Func1(a string) {

}

func Func2(a int) string {
	return "hello"
}

func Func3(a int) (string, string) {
	return "hello", "World"
}

func Func4(a int) (name string, age string) {
	return "hello", "World"
}

func Func5(num int) (name string, age int) {
	name = "Daming"
	age = 18
	return
}

func Func6(nums ...int) (name string, age int) {
	return
}

func Func7() {
	print("Functional 7")
}

func Func8() {
	myFunc := Func7
	myFunc()
}

func Func9() {
	fn := func(age int) string {
		return "I am " + strconv.Itoa(age) + " years old."
	}

	println(fn(3))
}

func functional7() func() string {
	return func() string {
		return "Oliver"
	}
}

func Func10() func(name string) string {
	return func(name string) string {
		return "hello, " + name
	}
}

func DeferClosure1() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosure2() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

func DeferReturn1() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturn2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturn3() *MyStruct {
	a := &MyStruct{
		name: "Jerry",
	}
	defer func() {
		a.name = "Tom"
	}()
	return a

}

type MyStruct struct {
	name string
}

func DeferLoop1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("The address of i is %p, value is %d\n", &i, i)
		}()
	}
}

func DeferLoop2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferLoop3() {
	//var j int
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("The address of i is %p, value is %d\n", &j, j)
		}()
	}
}
