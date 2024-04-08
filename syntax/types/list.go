package main

type List interface {
	Add(index int, val any) error
	Append(val any) error
	Delete(index int) error
}
