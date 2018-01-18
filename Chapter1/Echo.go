package main

import (
	"fmt"
	"os"
	"time"
	"strings"
	"strconv"
)


func main() {
	//打印args[0]
	fmt.Println("os.args[0]:", os.Args[0])

	s, sep := "", ""
	for _, input := range os.Args[1:] {
		s += sep + input;
		sep = " ";
	}
	fmt.Println("os.args:", s);

	//打印idx和input
	for idx, input := range os.Args[1:] {
		fmt.Println("os.rags.idx:", idx, "os.arg:", input);
	}

	var stringSlice [100000]string
	for i:=0;i<10000;i++{
		stringSlice[i]=strconv.Itoa(i)
	}

	//测试join 和loop时间
	currentTime := time.Now().Nanosecond()
	for _, input := range stringSlice{
		s += sep + input;
		sep = " ";
	}
	fmt.Println("loop spent ", time.Now().Nanosecond()-currentTime)

	currentTime = time.Now().Nanosecond()
	strings.Join(stringSlice[0:], sep)
	fmt.Println("join spent ", time.Now().Nanosecond()-currentTime)
}
