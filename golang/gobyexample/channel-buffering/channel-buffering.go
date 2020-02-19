package main

import "fmt"

func main()  {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}


// messages <- "channel"をコメントアウト
//→
// buffered
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /Users/okamotomasahiro/19_summer_intern/transcribe/transcribe/golang/gobyexample/channel-buffering/channel-buffering.go:12 +0x10b
