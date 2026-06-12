// select
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c { //防止没数了还在读
		time.Sleep(time.Second)
		fmt.Printf("wroker %d receiver %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c) //要创一个协程收
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second) //计时
	tick := time.Tick(time.Second)     //每隔一段时间送个值过来
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select { //防止阻塞,生成数据和收数据的速度是不一样的
		case n := <-c1: //可以使用nil channel，数据没准备好也可以写
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //间隔时间太长就会来这里
			fmt.Println("timeout")
		case <-tick: //定时输出队长
			fmt.Println("queue len=", len(values))
		case <-tm: //获得数据后就return掉
			fmt.Println("bye")
			return

		}
	}
}
