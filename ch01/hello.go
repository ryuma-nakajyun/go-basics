package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

// 定数宣言
const TimeFormatMilli = "2006-01-02 15:04:05.000"

// readFile
func readFile(name string) {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println(currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	f, _ := os.Open(name)
	defer f.Close()

	s := bufio.NewScanner(f)
	fmt.Println(currentTime.Format(TimeFormatMilli), "都道府県情報を出力")
	for s.Scan() {
		fmt.Println(s.Text())
	}

	currentTime = time.Now()
	fmt.Println(currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}

// main
func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println(currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	// time.Sleep(1 * time.Second) // 1秒だけ生かす

	readFile("pref.csv")

	currentTime = time.Now()
	fmt.Println(currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
