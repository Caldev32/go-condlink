package main

import(
	"fmt"
	"time"
	"net"
)

func Check(destination string, port string) string{
	endereco:= destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp",endereco, timeout)
	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v é inacessivel, \n Error: %v",destination, err)

	}else{
		status = fmt.Sprintf("[UP] %v é acessivel, \n De: %v\n Para %v", destination,conn.LocalAddr(), conn.RemoteAddr())

	}
	return status
}