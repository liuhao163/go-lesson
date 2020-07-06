package recover

import "fmt"

func Recovery() {
    for i := 0; i < 10; i++ {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("捕获到的错误：%s\n", r)
            }
        }()
    
        if i%2 == 0 {
            panic(fmt.Sprintf("偶数退出：%v",i))
        }else{
            fmt.Println(i)
        }
    }
}

func process(i int){
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("捕获到的错误：%s\n", r)
        }
    }()
    
    if i%2 == 0 {
        panic(fmt.Sprintf("偶数退出：%v",i))
    }else{
        fmt.Println(i)
    }
}