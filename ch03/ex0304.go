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
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)), "func", fn)
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// end
	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}
