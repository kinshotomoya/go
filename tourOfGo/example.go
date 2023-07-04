package main

import "fmt"

func main() {
	basicType()
	typeConversion()
}

func basicType() {
	var (
		i1 int = 1 // 符号あり。32bitか64bit（CPUに依存）
		i2 int16 = -10 // 符号あり
		i3 int32 = -100 // 符号あり
		i4 int64 = -1000 // 符号あり　
		i5 uint = 1 // 符号なし
		i6 uint8 = 2 // 符号なし
		i7 rune = 333 // int32と同じ。unicodeのコードポインタとして利用される
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