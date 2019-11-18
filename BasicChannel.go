package main

import "fmt"

func main() {
	// this is how we create the channel make -> followed by "chan" and then type of the datatype we want to receive ex: string
	ch := make(chan string, 1)
	ch <- "hello"
	// this is how we receive the message from channel
	fmt.Println(<-ch)

}
