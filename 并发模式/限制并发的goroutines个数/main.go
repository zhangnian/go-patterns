package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func producer() chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			out <- i
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// 用于控制并发的goroutine数
	ch := make(chan bool, 10)

	taskQueue := producer()
	for task := range taskQueue {
		ch <- true

		go func(task int, ch chan bool) {
			defer func() {
				<-ch
			}()

			log.Printf("handle: %d\n", task)
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}(task, ch)
	}

	log.Println("main exit")
}
