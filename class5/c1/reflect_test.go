/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午9:29
* */
package c1

import (
	"fmt"
	"reflect"
	"testing"
)

type ren struct {
	Name string
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	r := ren{Name:"sdsd"}
	CheckType(f)
	CheckType(r)
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t)
	fmt.Println()
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown",t)
		of := reflect.ValueOf(v)
		fmt.Println(of)
	}
}
