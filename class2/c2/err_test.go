/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午4:13
* */
package c2

import (
	"fmt"
	"testing"
)

func division(a, b int) int {
	if b == 0 {
		panic("b = 0")
	}
	c := int(a / b)
	return c
}

func TestDivision(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	i := division(1, 0)
	fmt.Println(i)
}
