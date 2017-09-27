package main

import (
	"log"
	"time"
)

func foo(){
	time.Sleep(time.Second * 3)
	log.Println("foo")
}

// 这是一个装饰器，用于计算函数的耗时
// 装饰器可以在不修改原有函数实现代码的基础上为其增加新的功能
func timeit(fn func()) func(){
	wrapper := func(){
		begin :=time.Now()

		// 调用被装饰的函数
		fn()
		cost := time.Since(begin)
		log.Printf("cost: %.2fs\n", cost.Seconds())
	}

	// 返回装饰函数
	return wrapper
}

func main(){
	foo := timeit(foo)
	// foo函数此时已经指向了装饰函数wrapper，而不是原来的那个foo函数了
	foo()
}
