package main

import (
	"io/ioutil"
	"fmt"
)

// switch后面可以没有表达式，在case里加条件
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "basic.go"
	// 条件里赋值的contents err 作用域只在if语句块
	// if条件不需要括号，可以赋值
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(83),
		grade(92),
		grade(100),
		//grade(101),
	)

}
