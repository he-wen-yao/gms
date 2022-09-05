package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	_   A
	age int
}

func main() {
	b := &B{age: 1, Name: "你好"}
	fmt.Println(b)
}
