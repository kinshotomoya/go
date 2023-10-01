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

// ↓処理の流れ
//1. main関数が呼び出されると、main関数のための新しいスタックフレームが生成されます。
//2. このスタックフレーム内には、wg (sync.WaitGroupのインスタンス) と numbers (intのスライス) が含まれます。
//3. forループが実行されると、index変数もこのスタックフレーム内に格納されます。
//4. go func()が呼び出されるたびに、新しいゴルーチンが作成され、それに伴って新しいスタックが割り当てられます。この匿名関数のクロージャがindexをキャプチャするため、indexの実際の値がヒープ上に格納される可能性があります。
//5. それぞれのゴルーチンは、それぞれのスタック上に自分のスタックフレームを持ちます。匿名関数がキャプチャした変数への参照は、このゴルーチンのスタックフレーム上にも存在しますが、実際のindexの値はヒープ上にあります。
//6. この動作の結果として、すべてのゴルーチンは同じindexの値（正確には、ヒープ上の同じ位置を参照するポインタ）を共有します。そのため、最後のゴルーチンが実際に実行される前に、forループが完了しindexが9になる場合、すべてのゴルーチンがこの9の値を見る可能性があります。
//
//この説明は概念的なもので、実際のメモリの配置や最適化の詳細はGoの実装や使用されるコンパイラに依存します。しかし、この説明はコードの動作を理解するのに役立つ基本的な概念を提供します。
// indexがヒープにmoveしており、各goroutineが保持するスタックフレームではヒープに存在するindex値への参照ポイントを格納している
// なので同じindex値を参照することになり、goクロージャで表示される値が同じ値になる可能性があるということ
