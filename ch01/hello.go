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
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime, "start:", runtime.FuncForPC(pc).Name())

	fmt.Println("Hello, world")
	time.Sleep(1 * time.Second) // 1秒だけ生かす

	readFile("sample.txt")
	
	currentTime = time.Now()
	fmt.Println("currentTime", currentTime, "end:", runtime.FuncForPC(pc).Name())
}


