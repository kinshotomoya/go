package main

import "sync"

func main() {
	var wg sync.WaitGroup
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// ↓indexがfor文を抜けても利用できる（goroutineから参照される）のはヒープに移動しているから
	// go build -gcflags '-m' for.go
	// ↓結果
	//[git][* main]:~/work_space/go/tourOfGo/ go build -gcflags '-m' for.go                                                                                        [/Users/jinzhengpengye/work_space/go/tourOfGo]
	//# command-line-arguments
	//./for.go:21:6: can inline Print
	//./for.go:10:6: can inline main.func1
	//./for.go:11:11: inlining call to sync.(*WaitGroup).Done
	//./for.go:12:9: inlining call to Print
	//./for.go:6:6: moved to heap: wg
	//./for.go:8:6: moved to heap: index
	//./for.go:7:18: []int{...} does not escape
	//./for.go:10:6: func literal escapes to heap
	for index := range numbers {
		wg.Add(1)
		go func() {
			wg.Done()
			Print(index)
		}()
	}
	print("loop finished")
	wg.Wait()
	print("finished")

}

func Print(number int) {
	print(number)

}
