package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "string"
	messages <- "buffered"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
