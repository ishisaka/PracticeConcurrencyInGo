// Channel
package main

import (
	"fmt"
)

func main() {
	instream := make(chan int)
	go func() {
		defer close(instream)
		for i := 1; i <= 5; i++ {
			instream <- i
		}
	}()

	for integer := range instream {
		fmt.Printf("%v ", integer)
	}
	fmt.Println()
}
