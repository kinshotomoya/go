package main

import (
	"fmt"
)

func main() {
	pointer()
	structs()
	arrays()
	slices()
}

func pointer() {
	// 続き
	// https://go-tour-jp.appspot.com/moretypes/1
	i := 1
	p := &i

	fmt.Println(p)  // メモリアドレス
	fmt.Println(*p) // 実体

	*p = 100
	fmt.Println(p)  // メモリアドレス
	fmt.Println(*p) // 実体

}

// NOTE: goのメモリ管理に関しては↓参照
//https://zenn.dev/kazu1029/articles/38ab3d99ef0de3

func structs() {

	type Person struct {
		name string
		age  int
	}

	person := Person{"javen", 1}
	fmt.Println(person.name)

	person2 := Person{"javen", 100}
	p := &person2
	fmt.Println(p)

}

func arrays() {
	var a [2]string
	a[0] = "a"
	a[1] = "b"

	fmt.Println(a)

	// arrayは固定長なので、[1 2 4 5 0 0]のように余白には0がはいる
	primes := [6]int{1, 2, 4, 5}
	fmt.Println(primes)

	names := [3]string{"", ""}
	fmt.Println(names)
}

func slices() {
	// 可変長
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	// arrayからsliceを作成できる
	arrays := [6]int{1, 2, 3, 4, 5, 6}
	slice := arrays[1:3]
	fmt.Println(slice)

	// sliceはarrayの参照である
	// sliceのデータ構造として、配列へのポインタを保持している
	// https://www.geeksforgeeks.org/slices-in-golang/
	// ↓の配列はコンパイル時にサイズがわかり、スタックの最大サイズを超えないのでスタックに保存される
	names := [4]string{"kinsho", "tanaka", "tanabe", "ishikawa"}
	// このスライスもサイズがわかっているので、スタックに保存
	names1 := names[0:2]
	names2 := names[1:3]

	fmt.Println(names1)
	fmt.Println(names2)

	// ↓　元のarrayの値が変更される
	names1[1] = "kimura"

	// ので、どちらのスライスも値が変更されている
	fmt.Println(names1)
	fmt.Println(names2)

}
