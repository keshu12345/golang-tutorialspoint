package main

import "fmt"

func main() {

	userChan := make(chan string)
	userChan <- "Keshav Rajput" // blocking the sender channel

	userVal := <-userChan
	fmt.Println(userVal)
}
