package main

import (
	"time"
	"fmt"
)

func main(){
	t := time.Now()
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(t.Weekday(), t.Year(), t.Month(), t.Day())
}
