package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// 定数宣言
const TimeFormatMilli = "2006-01-02 15:04:05.000"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

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

func main() {
	logger := slog.New(newHandler())

	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	// start
	currentTime := time.Now()
	logger.Info(
		fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)),
		"func", fn,
	)

	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"
	fmt.Println(x)
	fmt.Println(y)

	// end
	currentTime = time.Now()
	logger.Info(
		fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)),
		"func", fn,
	)
}
