/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午5:07
* */
package c3

import (
	"context"
	"fmt"
	"testing"
)

func TestService(t *testing.T) {
	dataCh := make(chan int, 10)
	end := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	go production(dataCh, cancel)
	go consume(ctx, dataCh, end)
	<-end
}

// 生产者
func production(ch chan int, cancelFunc context.CancelFunc) {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	cancelFunc()
}

// 消费者
func consume(ctx context.Context, ch chan int, end chan bool) {
forloop:
	for {
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-ctx.Done():
			fmt.Println("end------------------")
			break forloop
		}
	}
	end <- true
}
