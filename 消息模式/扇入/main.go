package main

import (
	"sync"
	"log"
	"time"
)

// 扇入模式，从多个channel中接受数据，并汇集到一个channel中
func FanIn(in_channels ...<-chan int) <-chan int {
	out_channel := make(chan int)

	var wg sync.WaitGroup

	for _, channel := range in_channels {
		wg.Add(1)
		go func(channel <-chan int) {
			defer wg.Done()
			for item := range channel {
				out_channel <- item
			}
		}(channel)
	}

	go func(){
		wg.Wait()
		close(out_channel)
	}()

	return out_channel
}


func main() {
	ch_in1 := make(chan int)
	ch_in2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=0; i < 10; i++{
			ch_in1 <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch_in1)

	}()

	wg.Add(1)
	go func(){
		wg.Done()
		for i:=10; i < 20; i++{
			ch_in2 <- i
			time.Sleep(time.Millisecond * 400)
		}
		close(ch_in2)
	}()


	wg.Add(1)
	out_ch := FanIn(ch_in1, ch_in2)
	go func(){
		defer wg.Done()
		for i:= range out_ch{
			log.Println(i)
		}
	}()

	wg.Wait()
	log.Println("main goroutine exit")
}
