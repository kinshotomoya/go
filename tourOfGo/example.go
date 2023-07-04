package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	basicType()
	typeConversion()
	typeInterface()
	constant()
	constantNumber()
	forstate()
	for2()
	for3()
	ifWithShortStat()
	swhich()
	switch2()
	defers()
}

func basicType() {
	var (
		i1 int     = 1     // 符号あり。32bitか64bit（CPUに依存）
		i2 int16   = -10   // 符号あり
		i3 int32   = -100  // 符号あり
		i4 int64   = -1000 // 符号あり
		i5 uint    = 1     // 符号なし
		i6 uint8   = 2     // 符号なし
		i7 rune    = 333   // int32と同じ。unicodeのコードポインタとして利用される
		i8 float32 = 1.1
		i9 float64 = 1.1
	)
	fmt.Printf("%T", i1, i2, i3, i4, i5, i6, i7, i8, i9)
}

func typeConversion() {
	var x, y int = 3, 4

	// 明示的に型変換できる
	f := float64(x)
	u := uint(y)

	fmt.Printf("\n%T", f, u)

}

func typeInterface() {
	// 明示的に型を指定しなくても型推論する
	v := true
	fmt.Printf("\nv is type of %T", v)
}

func constant() {
	// 定数を
	const Name = "javen"
	fmt.Println("\n", Name)
}

func constantNumber() {
	// シフト計算
	// ↓はint型
	const Big = 1 << 60
	fmt.Printf("type is %T. value is %v", Big, Big)

	const Small float64 = 1 << 100
	// 浮動小数点
	// 指数部・仮数部の2進数で表現することで、int64よりも大きな数を表現できる
	fmt.Printf("\ntype is %T. value is %v", Small, Small)
}

func forstate() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println("\n", sum)
}

func for2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("\n", sum)
}

func for3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("\n", sum)
}

func ifWithShortStat() {
	var x, y, lim float64 = 3, 2, 10
	// if文に変数を定義できる
	// これはifブロック内でのみ利用できる
	if v := math.Pow(x, y); v < lim {
		fmt.Println(v)
	} else {
		// else内でも使える
		fmt.Println(v)
	}
}

func swhich() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch2() {
	currentDay := time.Now()
	switch day := currentDay.Weekday(); day {
	case time.Sunday:
		fmt.Println("wooooo")
	case time.Monday:
		fmt.Println("buuuuu")
	default:
		fmt.Println("nothing")
	}
}

func defers() {
	word := "hello"
	// deferは即時評価されるが、実行はreturnの前
	defer fmt.Println(word)

	word = "hiiii"

	fmt.Println("ssss")
}
