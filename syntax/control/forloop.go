package main

import "fmt"

func ForMap() {
	m := map[string]int{
		"1": 100,
		"a": 102,
	}

	for k, v := range m {
		println(k, v)
	}

	for k := range m {
		println(k, m[k])
	}
}

func LoopBug() {
	users := []User{
		{
			Name: "Tom",
		},
		{
			Name: "Jerry",
		},
	}

	m := make(map[string]*User, 2)
	for _, u := range users {
		m[u.Name] = &u
	}

	for k, v := range m {
		fmt.Printf("name: %s, user: %v\n", k, v)
	}
}

type User struct {
	Name string
}
