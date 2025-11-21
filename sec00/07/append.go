package main

import (
	"fmt"
	// "sec00/00"
	// "strconv"
)

func main(){
	a := [3]int{10, 20, 30}
	b := a[0:2]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b,100)
	fmt.Println(b)
	b = append(b, 1000)
	fmt.Println(b)
}
