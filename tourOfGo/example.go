package main

import "fmt"

func main() {
	basicType()
	typeConversion()
	typeInterface()
	constant()
	constantNumber()
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

	const Small float64 = 1 << 60
	// TODO: 浮動小数点の理解
	fmt.Printf("\ntype is %T. value is %v", Small, Small)
}