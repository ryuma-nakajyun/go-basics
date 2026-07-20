package main

import "fmt"

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

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
		var x = [...]int{1 ,2, 3}
		var y = [3] int{1, 2, 3}
		fmt.Println(x == y)
	}

	fmt.Println("===== 多次元配列のシミュレート =====")
	{
		var x = [2][3]int
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

		var y = [2][3]int
		fmt.Println("len(y):", len((y))
	}

	fmt.Println("===== 3.2　スライス =====")

	fmt.Println("===== 3.3 文字列、rune、バイト=====")

	fmt.Println("===== 3.4　マップ =====")
	
	fmt.Println("===== 3.5　構造体 =====")

		currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}

