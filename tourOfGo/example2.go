package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	pointer()
	structs()
	arrays()
	slices()
	slices2()
	slices3()
	makeSlice()
	jsonEncode()
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

func slices2() {
	s := []struct {
		i int
		b bool
	}{
		{
			1,
			true,
		},
	}
	fmt.Println(s)

	n := []int{1, 3, 5, 7, 11}

	n = n[:3]
	fmt.Println(n)

	n = n[2:]
	fmt.Println(n)

}

func slices3() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = s[:4]

	fmt.Printf("len: %d, capacity: %d\n", len(s), cap(s))
	s = append(s, 1, 1, 1, 1, 1, 1, 1)
	fmt.Printf("len: %d, capacity: %d\n", len(s), cap(s))

	var nillS []int
	fmt.Println(nillS)

}

func makeSlice() {
	// makeを利用してsliceを作成することで、事前にsliceの要素数（length）、データ元のarrayの要素数（capacity）を指定して
	// sliceを作成することができる
	// こうすることで、データ元の固定長であるarrayのサイズを超えたらさらにメモリア割当てしてarrayを拡張するというパフォーマンスが
	// 悪い処理をすることになってしまう
	// arrayのサイズ拡張は、元の2倍ずつ拡張されるっぽい（2 -> 4 -> 8 -> 16）
	// 参考：https://zenn.dev/rescuenow/articles/41b02c28a9b06f
	s := make([]int, 5)
	fmt.Printf("len: %d, capacity: %d\n", len(s), cap(s))

	// TODO: 長さ指定しない場合（[]int{}でslice作成）と長さ指定する場合（make([]int, n)）とでappendにどれくらい差があるのかやってみる
}

func jsonEncode() {
	body := map[string]string{
		"name": "javen",
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		fmt.Println("JSON marshal error: ", err)
		return
	}

	fmt.Println(string(bytes))
}
