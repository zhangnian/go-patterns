package main

import (
	"log"
	"sync"
	"time"
)

type Limiter struct {
	curTime time.Time
	max     int
	current int
	sync.Mutex
}

func NewLimiter(max int) *Limiter {
	return &Limiter{
		curTime: time.Now(),
		max:     max,
	}
}

func (this *Limiter) IsLimited() (limit bool) {
	this.Lock()
	defer this.Unlock()

	limit = false
	if this.current >= this.max {
		if time.Since(this.curTime) > time.Second {
			this.curTime = time.Now()
			this.current = 1
		} else {
			limit = true
			return
		}
	} else {
		this.current += 1
	}

	return
}

func main() {
    // 限制每秒只能处理3个请求，多余的请求将会直接丢弃
	limiter := NewLimiter(3)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			if !limiter.IsLimited() {
				log.Printf("got %d\n", i)
			}

		}(i)

	}

	wg.Wait()
}
