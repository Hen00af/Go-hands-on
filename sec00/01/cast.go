package main

import (
	"fmt"
)

func main(){
	//型変換はエラーの温床になるという理念から、Go言語では明示的な型変換が必須。
	var a int32 = 42
	var b int64 = int64(a)
	var c float32 = float32(b)

	fmt.Println("int32 a:", a)
	fmt.Println("int64 b:", b)
	fmt.Println("float32 c:", c)
}
