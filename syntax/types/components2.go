package main

import "fmt"

// Person struct represents a general person
type Person struct {
	Name string
	Age  int
}

// Greet allows a Person to introduce themselves
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Student embeds Person. It "inherits" Person's fields and methods.
type Student struct {
	Person // Embedded struct
	School string
	Grade  int
}

func main() {
	// Initialize a Student
	student := Student{
		Person: Person{
			Name: "Alice",
			Age:  20,
		},
		School: "University of Go",
		Grade:  2,
	}

	// Access fields from the embedded Person struct directly on Student
	fmt.Println(student.Name) // Output: Alice
	fmt.Println(student.Age)  // Output: 20

	// Call a method defined on the embedded Person struct
	fmt.Println(student.Greet()) // Output: Hello, my name is Alice

	// Access Student's own fields
	fmt.Println(student.School) // Output: University of Go
	fmt.Println(student.Grade)  // Output: 2
}
