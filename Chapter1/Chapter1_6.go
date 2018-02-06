package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { //todo  practice 1.8 添加Http
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url error:%v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)

		fmt.Println("========statuscode:%v=========", resp.StatusCode) //todo practice1.9 输出状态码
		//res, err := io.Copy(os.Stdout, resp.Body)  todo practice 1.7 用copy方法输出到标准输出
		//do homework 1
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s:%v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf(string(b))
	}
}

func Fetch(url string, ch *chan<- string) {

}
