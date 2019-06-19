/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-18
* Time: 下午9:00
* */
package c1

import (
	"encoding/json"
	"errors"
	"testing"
)

// 表格测试法
func TestSquare(t *testing.T) {
	inputs := [...]int{1,2,3} // 测试数据
	expected := [...]int{1,4,9} // 期望值
	for i:=0;i<len(inputs);i++ {
		if square(inputs[i]) != expected[i] {
			t.Fatal(errors.New("err"))
		}
	}
}

type datamodel struct {
	Host string `json:"host"`
}

func BenchmarkSquare(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		//测试代码
		//data :=`{"host":"123456"}`
		//d := &datamodel{}
		//json.Unmarshal([]byte(data),d)
		d := &datamodel{
			Host:"1212312",
		}
		json.Marshal(d)
	}
	b.StopTimer()
}

