package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"
)

const TimeFormatMilli = "2006-01-02 15:04:05.000"
const r0711world = "r0711world_utf8.tsv"

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

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		// line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// isPrefix が true の場合は「1行が長すぎて分割されている」
		// 必要なら結合処理を書く
		fmt.Println(string(line))
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

	readFile(r0711world)

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}
