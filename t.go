package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(len(os.Args))
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])
	//var args map[string]int
	args := make(map[string]int)
	var targs map[string]int
	fmt.Println(args == nil)
	fmt.Println(len(args) == 0)
	args["carol"] = 21
	age, ok := args["test"]
	fmt.Println(age, ok)
	fmt.Println(args["test"])
	fmt.Println(args)
	fmt.Println(equal(args, targs))
}

func equal(x, y map[string]int) bool  {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range(x) {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}