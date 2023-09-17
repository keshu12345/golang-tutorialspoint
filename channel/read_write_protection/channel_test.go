package main

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()

	for i := 0; i < 10; i++ {

		// go func(i int) {

		server.userChan <- fmt.Sprintf("user %d", i)
		// server.addUser(fmt.Sprintf("user %d", i))

		// }(i)
		// go server.addUser(fmt.Sprintf("user %d", i))
	}

	fmt.Println("loop is don!!!")
}
