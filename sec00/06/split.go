package main

import(
	"fmt"
	"strings"
	// hello "sec00/00"
	"os"
)

func main(){
	a := [5]int{10,20,30,40,50}
	b := a[0:3]
	c := os.Environ()
	d := strings.Split(c[0],"=")
    fmt.Println(a)
    fmt.Println(b)
	fmt.Println(c)

    // range はインデックス(i)と値(value)を返す
    for _, value := range d {
        fmt.Println(value)
    }
}
