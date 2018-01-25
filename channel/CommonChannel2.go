package main

import "fmt"

var a1 string
var c1 = make(chan string)

func f1() {
	a1 = "hello, world"
	c1 <- a1
}

func main() {
	go f1();
	a2 := <-c1
	fmt.Print(a2)
}
