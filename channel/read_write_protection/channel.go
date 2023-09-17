package main

import (
	"fmt"
	"time"
)

type Server struct {
	users    map[string]string
	userChan chan string
	quit     chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:    make(map[string]string),
		userChan: make(chan string),
		quit:     make(chan struct{}),
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func (s *Server) Start() {

	go s.loop()
}
func (s *Server) loop() {

	for {
		select {
		case msg := <-s.userChan:
			fmt.Println(msg)

		case <-s.quit:
			fmt.Println("server want to quit")
			return
		default:
		}

	}

	// free:
	// 	for {
	// 		select {
	// 		case msg := <-s.userChan:
	// 			fmt.Println(msg)

	// 		case <-s.quit:
	// 			break free
	// 		}

	// }
}

// func (s *Server) loop() {

// 	for {
// 		user := <-s.userChan
// 		s.users[user] = user
// 		fmt.Printf("Adding user %s\n", user)
// 	}
// }

func main() {

	server := NewServer()
	server.Start()

	for i := 0; i < 10; i++ {

		// go func(i int) {

		server.userChan <- fmt.Sprintf("user %d", i)
		// server.addUser(fmt.Sprintf("user %d", i))

		// }(i)
		// go server.addUser(fmt.Sprintf("user %d", i))
	}

	go func() {
		time.Sleep(1 * time.Millisecond)
		close(server.quit)
	}()

	// select {}

	// server.quit <- struct{}{}

	// // block forever
	// // select {}

}

func sendMessage(msgChan chan<- string) {

	msgChan <- "hello world!"
}

func readMessage(msgChan <-chan string) {

	msg := <-msgChan
	fmt.Println(msg)
}
