package main

import (
	"fmt"
	"os"
	"strings"
	"runtime"
)

var Pi float64

func init() {
	Pi := 3.14
}

func main() {
	const (
		a = iota
		b
		c
		d = iota+50
		e
	)
	f, g, h := 1, 2, 3
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(runtime.Version())
	fmt.Println(a,b,c,d,e,f,g,h)
	fmt.Println(Pi)
}
