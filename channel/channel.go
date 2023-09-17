package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	userChan := make(chan string)

	go func() {

		for i := 1; i <= 10; i++ {

			userChan <- fmt.Sprintf("user_%d", i)
		}
		close(userChan)
	}()

	for userVal := range userChan {

		fmt.Println(userVal)
	}

	fmt.Println(time.Since(start))

}
