package main

import (
	"fmt"
	"time"
)

func main() {
	userChan := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		// close(userChan)
		userChan <- "hello  Bob"
	}()

	for {
		select {
		case user, ok := <-userChan:

			if !ok {
				return
			} else {
				fmt.Println(user)
			}

		}
	}
	userName, ok := <-userChan

	if !ok {
		fmt.Println("User channel close returning error")
		return
	}

	if len(userName) == 0 {
		panic("uai ai ai user is not good")

	}

	fmt.Println(userName)

}
