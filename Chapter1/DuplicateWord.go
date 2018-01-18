package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"io/ioutil"
)

func main() {

	fileName := "/Users/didi/goworkspace/src/go-lesson/Chapter1/liuhao.txt"

	fmt.Println("====now searchDuplicateByFile by scanner======")
	file, err := os.Open(fileName);
	if err != nil && err != io.EOF {
		panic(err)
	}
	countMapFile := make(map[string]int,10)
	searchDuplicateByFile(countMapFile, file)

	fmt.Println("====now searchDuplicateByFile2 by ioutil ======")
	countMapFile2 := make(map[string]int,100)
	searchDuplicateByFile2(countMapFile2, fileName)

}

func searchDuplicateByFile(countMap map[string]int, f *os.File) {
	inputScanner := bufio.NewScanner(f)
	countWord(countMap, inputScanner)

	printDuplicate(countMap)
}

func searchDuplicateByFile2(countMap map[string]int, fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil && err != io.EOF {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line);
		words := strings.Split(line, " ")
		for _, word := range words {
			countMap[word]++ // 等价于 countMap[s]=conutMap[s]+1
		}
	}

	printDuplicate(countMap)

}

func countWord(countMap map[string]int, inputScanner *bufio.Scanner) {
	for inputScanner.Scan() {
		content := inputScanner.Text()
		content = strings.TrimSpace(content);
		words := strings.Split(content, " ")
		for _, word := range words {
			countMap[word]++ // 等价于 countMap[s]=conutMap[s]+1
		}
	}
}

func printDuplicate(countMap map[string]int) {
	for input, numbers := range countMap {
		if numbers > 1 {
			fmt.Printf(" duplicate key {%s}:\t{%d}", input, numbers)
			fmt.Println()
		}
	}
}
