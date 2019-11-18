package main

import (
	"fmt"
)

func main() {
	// initiallizing the channels
	msgch := make(chan Message, 1)
	errch := make(chan FailedMessage, 1)

	//define values to struct
	msg := Message{
		To:      []string{"sharath@sharath.com"},
		From:    "mohit@mohit.com",
		Content: "keep it secret, keep it safe",
	}

	failedMessage := FailedMessage{
		ErrorMessage:   "Message intercepted by the black rider",
		OrginalMessage: Message{},
	}

	//calling the channel and passing the value to the channels
	msgch <- msg
	errch <- failedMessage

	// select statement is like switch
	select {
	case receivedMessage := <-msgch:
		fmt.Print(receivedMessage)
	case receivedError := <-errch:
		fmt.Println(receivedError)
	default:
		fmt.Println("No message received from channels")

	}

}

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage   string
	OrginalMessage Message
}
