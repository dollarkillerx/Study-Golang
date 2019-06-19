/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午9:42
* */
package c2

import (
	"fmt"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name string `format:"normal,bs"`
	Age int
}

func (e *Employee) UpdateAge(val int) {
	e.Age = val
}


func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "DollarKiller", 20}
	// 按照名字获取成员
	//t.Logf("Name:value(%[1]v),Type(%[1]T)",reflect.ValueOf(*e).FieldByName("Name"))
	fmt.Println(reflect.ValueOf(*e).FieldByName("Name"))
	fmt.Println(reflect.ValueOf(*e).FieldByIndex([]int{2}))

	if field, err := reflect.TypeOf(*e).FieldByName("Name");err != true {
		t.Fatal(err)
	}else{
		format := field.Tag.Get("format")
		of := reflect.TypeOf(format)
		fmt.Println(of)
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	// MethodByName("UpdateAge") 参数名称
	// Call([]reflect.Value{reflect.ValueOf(1)})  1这个是参数

	t.Log(e)


}