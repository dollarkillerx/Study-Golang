/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午6:22
* */
package c8

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

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			task := runTask(i)
			ch <- task
		}(i)
	}
	allRet := ""
	//for i:=range ch {
	//	allRet += i + "\n"
	//}
	for i := 0; i < numOfRunner; i++ {
		allRet += <-ch + "\n"
	}
	return allRet
}

func TestAll(t *testing.T) {
	t.Log("Bef: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	//time.Sleep(time.Second)
	t.Log("End: ", runtime.NumGoroutine())
}
