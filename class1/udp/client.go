/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:19
* */
package main

import (
	"Study-Golang/err"
	"fmt"
	"net"
)

func main() {
	// 拨号获取连接
	conn, e := net.Dial("udp", "127.0.0.1:9001")
	err.CheckErr(e)
	defer conn.Close()

	for {
		fmt.Println("print: ")
		var data string
		fmt.Scan(&data)
		// 写出信息
		_, e = conn.Write([]byte(data))
		err.CheckErr(e)

		// 创建缓冲区
		buffer := make([]byte, 1024)

		// 从连接中获取信息
		n, e := conn.Read(buffer)
		err.CheckErr(e)
		fmt.Println("服务端返回: ", string(buffer[:n]))
	}
}
