package remake

import (
	"syscall"
)

func R_Putstr(s string) int{
	data := []byte(s) 
	l, err := syscall.Write(syscall.Stdout,data)
	if err != nil{
		return -1
	}
	return l
}
