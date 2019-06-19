/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-19
* Time: 上午10:28
* */
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"sync"
)

func main() {
	// cpu
	file, err := os.Create("cpu.prof")
	if err != nil {
		panic(err.Error())
	}
	err = pprof.StartCPUProfile(file)
	if err != nil {
		panic(err.Error())
	}
	defer pprof.StopCPUProfile()
	// cpu

	var wg sync.WaitGroup
	for i:=0;i<99;i++{
		wg.Add(1)
		go func() {
			ps()
			wg.Done()
		}()
	}

	// 内存
	mem, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(mem)
	
	// 多种tag
	//gos, _ := os.Create("gorutine.prof")
	//pprof.Lookup("")

	wg.Wait()

}

func ps() {
	num := rand.Intn(9999)
	pps(num)
}

func pps(s int) {
	for i:=0;i<s;i++{
		fmt.Println(i)
	}
}