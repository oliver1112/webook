package main

import "fmt"

type LinkedList struct {
	head *node

	Len int
}

func (l *LinkedList) Add(index int, val any) error {
	//TODO implement me
	fmt.Printf("%v\n", val)
	println(index)
	//println(val)
	return nil
}

func (l *LinkedList) Append(val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

type node struct {
	prev *node
	tail *node
}
