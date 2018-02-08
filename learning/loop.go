package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"bufio"
	"time"
)

// for的条件里不需要括号
// 可以省略初始条件、结束条件、递增表达式
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

// 每次从文件读取一行
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// 没有初始条件、没有递增条件，只有结束条件，分号省略，相当于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// while true的实现，死循环
func forever() {
	for {
		fmt.Println(time.Now(), "abc")
		time.Sleep(time.Second * 5)
	}
}

func main() {
	fmt.Println(
		convertToBin(5),       // 101
		convertToBin(13),      // 1101
		convertToBin(1092765), // 换行需要逗号
		convertToBin(0),
	)
	printFile("basic.go")
	forever()
}
