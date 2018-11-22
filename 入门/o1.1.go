package main

import (
	"fmt"
	"strings"
)

var a string
func main() {
	type TZ int
	var g TZ
	g = 10
	fmt.Println(g)

	a = "G"
	var a int = 5
	var b int32
	b = int32(a + a)
	fmt.Println(a, b)
	var c1 complex64 = 5 + 10i
	fmt.Println("The value is: %v", c1)
	j := `asdfcd\naaaa`
	fmt.Println(strings.HasPrefix(j, `b`))
	fmt.Println(strings.Contains(j, "aaa"))
	fmt.Println(strings.Index(j, "aaa"))
	fmt.Println(strings.Count(j, "a"))
	fmt.Println(strings.Repeat("a", 10))
	fmt.Println(j)
	f1()
}
func f1() {
	a = "O"
	fmt.Println(a)
	f2()
}
func f2() {
	fmt.Println(a)
}