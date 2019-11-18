package main

import "time"

func main() {

	godur, _ := time.ParseDuration("10ms")
	//appending go infront of function will make goroutine
	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()

	//just adding the delay to show the above 2 function executed
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
