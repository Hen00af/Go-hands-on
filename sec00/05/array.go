package main

import (
	"fmt"
	hello "sec00/00"
	"strconv"
	"strings"
)

func main(){
	x := hello.Input("input data")
	ar := strings.Split(x, " ")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, er := strconv.Atoi(ar[i])
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total: ", t)
err:
	fmt.Println("invalid input ")
}
