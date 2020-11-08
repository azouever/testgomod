package main

import "fmt"

var name = "xkx"

func hello(name string) string {
	return "hello, " + name
}

func echo(str string) string {
	return str
}

func main() {
	age := 25
	fmt.Printf("hello, %s ,%d .from vscode",name,age)
}
