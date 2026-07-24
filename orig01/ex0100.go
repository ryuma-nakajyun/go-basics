package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

const TimeFormatMilli = "2006-01-02 15:04:05.000"
const PrefFile = "pref.csv"

// slog の自動 timestamp を消すハンドラ
func newHandler() slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// slog が自動で付ける timestamp を削除
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
		AddSource: true, // file:line を出す
	})
}

// readFile
func readFile(name string) {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)), "func", fn)

	f, _ := os.Open(name)
	defer f.Close()

	fmt.Println("======= 都道府県情報を一括出力 =======")
	record := bufio.NewScanner(f)
	for record.Scan() {
		fmt.Println(record.Text())
	}

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}

// main
func main() {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)), "func", fn)

	readFile(PrefFile)

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}
