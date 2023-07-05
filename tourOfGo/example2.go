package main

import (
	"fmt"
)

func main() {
	pointer()
	structs()
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
