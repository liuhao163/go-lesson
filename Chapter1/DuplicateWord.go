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
	fmt.Println("====now searchDuplicateByFile by os.stdin======")

	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text()=="exit" {
			break
		}

		wordsArray := strings.Split(input.Text(), " ")
		for _, word := range wordsArray {
			fmt.Printf(word)
			counts[word]++
		}
	}

	for inputStr, idx := range counts {
		if idx > 1 {
			fmt.Printf("this is %s count is %d", inputStr, idx)
			fmt.Println()
		}
	}

	fileName := "/Users/didi/goworkspace/src/go-lesson/Chapter1/liuhao.txt"

	fmt.Println("====now searchDuplicateByFile by scanner======")
	file, err := os.Open(fileName);
	if err != nil && err != io.EOF {
		panic(err)
	}

	searchDuplicateByFile(file)

	fmt.Println("====now searchDuplicateByFile2 by ioutil ======")
	searchDuplicateByFile2(fileName)

}

func searchDuplicateByFile(f *os.File) {
	countMap := make(map[string]int, 10)
	inputScanner := bufio.NewScanner(f)
	countWord(countMap, inputScanner)

	printDuplicate(countMap)
}

func searchDuplicateByFile2(fileName string) {
	countMap := make(map[string]int, 100)

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
		content = strings.TrimSpace(content)
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
