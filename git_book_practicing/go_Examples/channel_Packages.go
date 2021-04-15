package main

import "fmt"

/*

Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and
receive those values into another goroutine.
*/
//create a new channel with make(chan val-type)
//send a value into a channel  using the channel <- syntax
//receives <-channel syntax

func main(){
	messages :=make(chan string)
	go func(){
		messages <-"ping"
	}()
	msg := <-messages
	fmt.Println(msg)

	// Channel Buffering  up to 2 values

	messages :=make(chan string,2)
	messages <-"buffered"
	messages <-"channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}