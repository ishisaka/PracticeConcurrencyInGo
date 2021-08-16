// ゴルーチンのリークを防ぐパターン
package main

import (
	"fmt"
	"time"
)

func main() {
	dowork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} { // ①
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Do somethings.
					fmt.Println(s)
				case <-done: // ②
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := dowork(done, nil)

	go func() { // ③
		// 1秒後に操作をキャンセルする。
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling dowork goroutine...")
		close(done)
	}()

	<-terminated // ④
	fmt.Println("Done.")
}

/*
❶ doneチャネルをdoWork関数に渡します。慣例として、このチャネルは第1引数にします。
❷ この行はどこにでもに存在するfor-selectパターンを使っています。case文の1つで
   doneチャネルからシグナルが送られたかどうかを確認しています。もし送られていた
   ら、ゴルーチンからreturnします。
❸ 1 秒以上経過したらdoWorkの中で生成されたゴルーチンをキャンセルする他のゴルー
   チンを生成します。
❹ ここでdoWorkから生成されたゴルーチンがメインゴルーチンとつながります。
*/
