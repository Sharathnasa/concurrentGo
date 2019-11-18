package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "Hello World from sharath"

	words := strings.Split(phrase, " ")

	//channels
	ch := make(chan string, len(words))

	for _, words := range words {
		ch <- words
	}

	// this is how we close the channel
	//close(ch)

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	//once the channel gets exhausted, we need to close the channel, otherwise deadlock error can occur
	// ranging over the channels
	for msg := range ch {
		fmt.Println(msg + " ")
	}

}
