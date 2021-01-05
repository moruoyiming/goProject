package main

import (
	"fmt"
)

var xx, yy int

var (
	a int
	b string
)

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!" + " GO Language !")
	var age int32 = 12
	var length float32 = 16.00

	var mm = what(4)

	a = 32
	b = "what the fck"

	xx := 3232
	yy := false

	bb, cc := what2(16, 20)

	const LENGTH int =19
	const WIDTH int = 29

	fmt.Println("what it is ? ", age, length, mm, a, b, xx, yy, bb, cc)
}

func what(value int) int {
	var x = 10
	var y = 5
	var area = x * y * value
	fmt.Println(" area of value ", area)
	return area
}

func what2(value int, value2 int) (int, int) {
	length := value * 10
	width := value2 * 20
	return length, width
}
