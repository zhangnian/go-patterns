package main

import (
	"sync"
	"fmt"
)

type SingletonMap map[string]string

var (
	instance SingletonMap
	once sync.Once
)

func GetInstance() SingletonMap{
	once.Do(func() {
		instance = make(SingletonMap)
	})
	return instance
}

func main(){
	var wg sync.WaitGroup
	for i:= 0; i < 10; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			ins := GetInstance()
			fmt.Printf("%p\n", ins)
		}()
	}

	wg.Wait()
}
