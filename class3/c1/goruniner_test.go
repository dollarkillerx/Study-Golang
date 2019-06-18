/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午4:27
* */
package c1

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	var cs sync.WaitGroup
	for i := 0; i < 1999; i++ {
		cs.Add(1)
		go func(i int) {
			fmt.Println(i)
			cs.Done()
		}(i)
	}
	cs.Wait()
}
