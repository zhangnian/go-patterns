package main

import (
	"fmt"
	"time"
	"sync"
)

type Pool struct {
	pool chan *Object
}

type Object struct {
	id   int
	pool *Pool
}

func NewObject(id int, pool *Pool) *Object {
	return &Object{id: id, pool: pool}
}

func (this *Object) Release() {
	this.pool.pool <- this
}

func NewPool(count int) *Pool {
	pool := &Pool{
		pool: make(chan *Object, count),
	}

	for i := 0; i < count; i++ {
		obj := NewObject(i, pool)
		pool.pool <- obj
	}

	return pool
}

func (this *Pool) Get() *Object {
	retry := 0
RETRY:
	for {
		select {
		case obj := <-this.pool:
			return obj
		default:
			if retry < 3{
				time.Sleep(time.Second)
				retry += 1
				continue RETRY
			}
			return nil
		}
	}
}

func main() {
	pool := NewPool(4)

	var wg sync.WaitGroup
	for i:= 0; i < 10; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			obj := pool.Get()
			fmt.Printf("%p\n", obj)
			time.Sleep(time.Millisecond * 900)
			obj.Release()
		}()
	}

	wg.Wait()
}
