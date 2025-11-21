package remake

import(
	"syscall"
)

func R_Cat (n string){
	fd, err := syscall.Open(n, syscall.O_RDONLY, 0)
	if err != nil{
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil{
			return
		}
		if n == 0{
			break
		}
		syscall.Write(syscall.Stdout, buf[:n])
	}
	syscall.Close(fd)
}
