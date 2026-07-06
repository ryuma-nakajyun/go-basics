package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	pc, _, _, _ := run\time.Caller(0)
	fmt.Println("start:" + runtime.FuncForPC(pc).Name())
	fmt.Println("Hello, world")
	time.Sleep(10 * time.Second) // 10秒だけ生かす
	fmt.Println("end:" + runtime.FuncForPC(pc).Name())
}

