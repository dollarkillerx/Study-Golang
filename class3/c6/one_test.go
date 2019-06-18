/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午6:04
* */
package c6

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func TestOnce(t *testing.T) {

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			once.Do(hello)
			wg.Done()
		}()
	}
	wg.Wait()
}

func hello() {
	fmt.Println("Hello")
}
