/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午4:40
* */
package c2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	nowTime := time.Now().UnixNano()
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
	tTime := time.Now().UnixNano()
	fmt.Println("time: ", (tTime - nowTime))
}

func TestMutCounter(t *testing.T) {
	nowTime := time.Now().UnixNano()
	var nut sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			nut.Lock()
			defer nut.Unlock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
	tTime := time.Now().UnixNano()
	fmt.Println("time: ", (tTime - nowTime))
}

func TestRWCounter(t *testing.T) {
	nowTime := time.Now().UnixNano()
	var rw sync.RWMutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			rw.Lock()
			defer rw.Unlock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
	tTime := time.Now().UnixNano()
	fmt.Println("time: ", (tTime - nowTime))
}
