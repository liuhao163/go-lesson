package main

import "fmt"

var a1 string
var c1 = make(chan int)

func f1() {
	a1 = "hello, world"
	<-c1
}

func main(){
	go f1();
	c1<-0
	fmt.Print(a1)
}
