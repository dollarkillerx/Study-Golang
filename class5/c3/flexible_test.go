/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午10:17
* */
package c3

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// reflect.Ptr 指针
// reflect.Struct 结构体
func fillBySettings(st interface{},settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() == reflect.Ptr { // 如果不是指针
		fmt.Println("ptr")
		// Elem() 获取指针指向的值
		return errors.New("st not ptr")
	}else if reflect.TypeOf(st).Elem().Kind() != reflect.Struct{ // 如果不是结构
		return errors.New("st not struct")
	}

	if settings == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k,v := range settings {
		if field, ok = reflect.TypeOf(st).Elem().FieldByName(k);!ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v))
		}

	}

	return nil
}

type hero struct {
	Name string `hr:"ho"`
}


// 基础测试
func TestBBst(t *testing.T) {
	h := &hero{}
	i := map[string]interface{}{
		"Name": "dollarkiller",
	}
	bs := fillBySettings(h, i)
	if bs != nil {
		t.Log(bs.Error())
	}else{
		fmt.Println(i)
	}
}
