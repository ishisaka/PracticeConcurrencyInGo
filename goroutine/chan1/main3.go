package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutbuff bytes.Buffer
	defer stdoutbuff.WriteTo(os.Stdout)

	intstream := make(chan int, 4)
	go func() {
		defer close(intstream)
		defer fmt.Fprintf(&stdoutbuff, "Procedure Done...\n")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutbuff, "Sending: %d\n", i)
			intstream <- i
		}
	}()

	for integer := range intstream {
		fmt.Fprintf(&stdoutbuff, "Received %v. \n", integer)
	}
}
