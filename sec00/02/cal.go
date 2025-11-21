package main

import (
	"fmt"
	"sec00/00"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("invalid number")
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
