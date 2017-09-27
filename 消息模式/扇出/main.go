package main

import (
	"log"
	"sync"
)

func FanOut(in <-chan int, len int) []chan int {
	ch_outs := make([]chan int, 0)
	for i := 0; i < len; i++ {
		ch_outs = append(ch_outs, make(chan int))
	}

	go func() {
		defer func() {
			for _, ch_out := range ch_outs {
				close(ch_out)
			}
		}()

		// 使用round-bin的方式将从input channel中获取到的数据写入多个output channel中
		selected_idx := 0
		for {
			ch_out := ch_outs[selected_idx]
			selected_idx = (selected_idx + 1) % len

			item, is_close := <-in
			if !is_close {
				return
			}
			ch_out <- item
		}
	}()

	return ch_outs
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	ch_in := make(chan int)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch_in <- i
		}

		close(ch_in)
	}()

	ch_outs := FanOut(ch_in, 10)

	// 每个outpout channel启动不同的goroutine去处理
	for idx, ch_out := range ch_outs {
		wg.Add(1)
		go func(idx int, ch_out <-chan int) {
			defer wg.Done()
			for item := range ch_out {
				log.Printf("groutine: %d process: %d\n", idx, item)
			}
		}(idx, ch_out)
	}

	wg.Wait()
	log.Println("main groutine exit")
}
