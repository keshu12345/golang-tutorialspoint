package main

import (
	"fmt"
	"time"
)

func main() {

	// userChan := make(chan string)
	// userChan <- "Keshav Rajput" // blocking the sender channel

	// userVal := <-userChan
	// fmt.Println(userVal)

	start := time.Now()

	userChan := make(chan string)

	for i := 1; i <= 10; i++ {

		userChan <- fmt.Sprintf("user_%d", i)
	}

	go func() {
		for userVal := range userChan {

			fmt.Println(userVal)
		}
	}()
	close(userChan)

	fmt.Println(time.Since(start))
}
