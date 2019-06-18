/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:38
* */
package main

import (
	"Study-Golang/err"
	"fmt"
	"net"
)

func main() {
	conn, e := net.DialUDP("udp", nil,
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 0, 255),
			Port: 9000,
		}) // 协议, 发送者,接收者
	defer conn.Close()
	err.CheckErr(e)

	for {
		fmt.Print("input: ")
		var data string
		fmt.Scan(&data)
		_, e := conn.Write([]byte(data))
		if e != nil {
			fmt.Println(e.Error())
			continue
		}
		fmt.Println("发送成功")
	}
}
