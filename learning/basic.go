package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

// 包内部变量
// 集中定义
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	// 变量zero值
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	// 定义多个变量
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 由编译器决定类型
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// short variable定义
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	// 复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// 必须强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func enums() {
	// 枚举类型定义
	const (
		cpp        = iota
		_           // 跳过
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func consts() {
	// 常量，Go语言中常量一般不会大写
	const filename string = "abc.txt"
	// const数值可以作为各种类型使用
	const a, b = 3, 4
	var c int
	// a/b不再需要强制转换
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func main() {
	fmt.Printf("Hello, world.\n")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	// 包变量
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
