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
	"sync"
	"testing"
)

func TestService(t *testing.T) {
	dataCh := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go production(dataCh, &wg)
	wg.Add(1)
	go consume(dataCh, &wg)

	wg.Wait()
}

// 生产者
func production(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

// 消费者
func consume(ch chan int, wg *sync.WaitGroup) {
forleep:
	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println(data)
			} else {
				break forleep
			}
		}
	}
	wg.Done()
}
