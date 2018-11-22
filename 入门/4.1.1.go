package main

import "fmt"

func main() {
	var a = []int{1,2,3,4,5,6,7,8,9}
	var b = make([]int, 10)
	b[0] = 1
	b[1] = 2
	a = append(a, 10)
	fmt.Println(sum(a[1:]))
}

func sum(a[] int) int{
	var sum int
	for _, v := range(a) {
		sum += v
	}
	return sum
}