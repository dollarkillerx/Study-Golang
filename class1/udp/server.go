/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:07
* */
package main

import (
	"Study-Golang/err"
	"fmt"
	"net"
)

func main() {
	// 解析得到UDP地址
	addr, e := net.ResolveUDPAddr("udp", ":9001")
	err.CheckErr(e)

	// 在UDP地址上建立UDP监听,得到连接
	conn, e := net.ListenUDP("udp", addr)
	err.CheckErr(e)
	defer conn.Close()

	// 建立缓冲区
	buffer := make([]byte, 1024)

	for {
		//从连接中读取内容,丢入缓冲区
		i, udpAddr, e := conn.ReadFromUDP(buffer)
		// 第一个是字节长度,第二个是udp的地址
		if e != nil {
			continue
		}
		fmt.Printf("来自%v,读到的内容是:%s\n", udpAddr, buffer[:i])

		// 向客户端返回消息
		conn.WriteToUDP([]byte("hello"), udpAddr)
	}
}
