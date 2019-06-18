/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午5:07
* */
package c3

import (
	"fmt"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	dataCh := make(chan int, 10)
	go production(dataCh)
	go consume(dataCh)
	time.Sleep(time.Second * 3)
}

// 生产者
func production(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
}

// 消费者
func consume(ch chan int) {
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		}
	}
}
