/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午6:14
* */
package c7

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", i)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			task := runTask(i)
			ch <- task
		}(i)
	}
	return <-ch
}

func TestFirst(t *testing.T) {
	fmt.Println("Go num: ", runtime.NumGoroutine())
	FirstResponse()
	time.Sleep(time.Second)
	fmt.Println("Go num: ", runtime.NumGoroutine())
}
