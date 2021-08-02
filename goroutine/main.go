// sync.WaitFroupの使い方
package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup
	// for _, salutation := range []string{"hello", "greetings", "good by"} {
	// 	wg.Add(1)
	// 	go func(salutation string) {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)
	// 	}(salutation)
	// }
	// wg.Wait()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("1st goroutine sleeping...")
	// 	time.Sleep(1)
	// }()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("2nd gorutine sleeping...")
	// 	time.Sleep(2)
	// }()
	// wg.Wait()
	// fmt.Println("All gorutine complete.")

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}

	const numGreaters = 5
	var wg sync.WaitGroup
	wg.Add(numGreaters)
	for i := 0; i < numGreaters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
