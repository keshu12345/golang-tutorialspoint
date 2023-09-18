package main

import (
	"fmt"
	"sync"
	"time"
)

func getUserFromDBD(userId int, ch chan int, wg *sync.WaitGroup) {

	time.Sleep(20 * time.Millisecond)

	ch <- userId
	wg.Done()

}

func getUserIdNetflixD(userId int, ch chan int, wg *sync.WaitGroup) {

	time.Sleep(100 * time.Millisecond)
	ch <- userId
	wg.Done()
}

func main() {

	now := time.Now().UTC()

	userIdsChan := make(chan int)

	wg := &sync.WaitGroup{}

	userId := 10

	wg.Add(2)
	go getUserFromDBD(userId, userIdsChan, wg)
	go getUserIdNetflixD(userId, userIdsChan, wg)

	wg.Wait()
	close(userIdsChan)

	for userIdFromChan := range userIdsChan {

		fmt.Println(userIdFromChan)

	}
	fmt.Println(time.Since(now))
}
