/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午3:59
* */
package c1

import (
	"errors"
	"fmt"
	"testing"
)

var (
	LessZeroError = errors.New("0错误")
)

func division(a, b int) (int, error) {
	if a >= 0 && b > 0 {
		c := int(a / b)
		return c, nil
	}
	return 0, LessZeroError
}

func TestDivision(t *testing.T) {
	//i, e := division(3, 4)
	i, e := division(3, 0)
	if e != nil {
		t.Fatal(e.Error())
	}
	fmt.Println(i)
}
