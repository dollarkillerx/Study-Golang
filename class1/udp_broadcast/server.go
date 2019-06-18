/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:31
* */
package main

import (
	"Study-Golang/err"
	"fmt"
	"log"
	"net"
)

// udp 广播
func main() {
	conn, e := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 9000,
	})
	defer conn.Close()
	err.CheckErr(e)

	buf := make([]byte, 1024)
	for {
		i, _, e := conn.ReadFromUDP(buf)
		if e != nil {
			log.Println(e.Error())
			continue
		}
		fmt.Println("data: ", string(buf[:i]))
	}
}
