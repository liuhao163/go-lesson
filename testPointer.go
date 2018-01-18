package main

import "fmt"

func main() {
	x := 10
	fmt.Println("地址", &x) //打印引用的地址
	pointer := &x         //声明一个指针
	pointer2 := &x //多指针指向一个地址
	fmt.Println("指针地址pointer", pointer)//打印下指针地址
	fmt.Println("指针的值pointer", *pointer)//打印下指针值
	fmt.Println("指针地址pointer2", pointer2)//打印下指针地址
	fmt.Println("指针的值pointer2", *pointer2)
	fmt.Println("==========================================")
	fmt.Println("传递引用之前的 x=", x)
	x1 := add(x)
	fmt.Println("传递引用之后的 x=", x)
	fmt.Println("传递引用之后的 x1=", x1)
	fmt.Println("传递引用之后的 x1=", x1)
	fmt.Println("==========================================")
	x2 := add1(&x)
	fmt.Println("传递引用之后的 x=", x)
	fmt.Println("传递引用之后的 x2=", x2)
	fmt.Println("==========================================")

	fmt.Println("传递之后看指针", *pointer2, *pointer, x2)
}
func add(i int) int {
	i = i + 1;
	return i;
}

func add1(i *int) int {
	fmt.Println("打印下参数的引用", i)
	*i = *i + 1
	return *i;
}
