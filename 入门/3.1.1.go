package main

import "fmt"

func main() {
	c := callb(6, Add)
	var arr = [10]string{3:"33333", 4:"44444"}
	fmt.Println(c, arr)
}

func Add(a, b int) int {
	return a + b
}

func callb(y int, f func(int, int) int) int {
	return f(y, 2)
}