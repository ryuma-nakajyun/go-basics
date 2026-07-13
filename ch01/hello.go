package main

import (
	"fmt"
	"time"
	"runtime"
	"os"
	"bufio"
)

func readFile(name string) {
	f, _ := os.Open(name)
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

func main() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("start:" + runtime.FuncForPC(pc).Name())
	fmt.Println("Hello, world")
	time.Sleep(3 * time.Second) // 3秒だけ生かす

	readFile("sample.txt")
		
	fmt.Println("end:" + runtime.FuncForPC(pc).Name())
}

