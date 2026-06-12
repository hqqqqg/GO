// 传统的 waitgroup等待
package main

import (
	"fmt"
	"sync"
)

// 给worker分配channel
func doworker(id int, w worker) {
	for n := range w.in { //防止没数了还在读
		fmt.Printf("wroker %d receiver %c\n", id, n)
		w.done()
		// //如果大小写分开等true的话就不用go开一个
		// go func() { //并发，不让打完小写后卡住大写
		// 	done <- true
		// }() //通知外面事情做完了
	}
}

type worker struct {
	in   chan int
	done func()
}

// 返回只能收的channel
func createrworker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() { //done是一个funtion
			wg.Done()
		},
	}
	go doworker(id, w) //打印
	return w

}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createrworker(i, &wg) //返回的结构
	}
	wg.Add(20) //加20个
	for i, worker := range workers {
		worker.in <- 'a' + i //发是阻塞式的，必须有收的

	}

	for i, worker := range workers {
		worker.in <- 'A' + i

	}
	wg.Wait()

}

func main() {
	chanDemo()

}
