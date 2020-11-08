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

// it's better only have one init function in one package 
func init(){
	println("import package will be executed only once")
}
