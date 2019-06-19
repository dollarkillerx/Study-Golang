/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午7:12
* */
package c10

import (
	"fmt"
	"sync"
	"testing"
)

func TestCcS(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	// 获取
	i := pool.Get().(int)
	fmt.Println(i)

	// 使用完后放回
	pool.Put(i)
}
