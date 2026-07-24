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

	fmt.Println("===== 3.1　配列 =====")
	{
		var x [3]int
		fmt.Println(x)
	}
	{
		var x = [3]int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	fmt.Println("===== 配列の比較 =====")
	{
		var x = [...]int{1, 2, 3}
		var y = [3]int{1, 2, 3}
		fmt.Println(x == y)
	}

	fmt.Println("===== 多次元配列のシミュレート =====")
	{
		var x = [2][3]int{}
		fmt.Println(x)
	}

	fmt.Println("===== インデックス（添字）の指定 =====")
	{
		var x [3]int
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])

		var y [2][3]int
		y[0][2] = 12
		y[1][1] = 3
	}

	fmt.Println("===== len =====")
	{
		var x [3]int
		fmt.Println("len(x):", len(x))
		var y = [2][3]int{}
		fmt.Println("len(y[0]):", len(y[0])) // 3
	}

	fmt.Println("===== 3.2　スライス =====")

	fmt.Println("===== 3.3 文字列、rune、バイト=====")

	fmt.Println("===== 3.4　マップ =====")

	fmt.Println("===== 3.5　構造体 =====")

	// end
	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}
