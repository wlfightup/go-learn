package main

import "fmt"

func main() {
	var i int
	defer test_defer()
	i = 7
	fmt.Println("%d -- %p", i, &i)
	var p *int
	p = &i
	fmt.Println("%d -- %p", *p, p)
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 7:
		fmt.Println(7)
	default:
		fmt.Println("error")
	}
	for j:=0; j<i; j++ {
		fmt.Print(j)
	}
	var d *int
	d = &i
	multinum(1, 2,3, d)
	fmt.Println(*d)
	fmt.Println("=========")
	fmt.Println(minisize(6, 1,-1,9,8))
}

func multinum(a int, b int, c int, d *int) (int, int) {
	i := 1000
	*d = i
	return a*b*c, c
}

func minisize(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func test_defer() {
	fmt.Println("------------------")
}