// 协程双向
// 论文communication sequential process
package main

import (
	"fmt"
	"time"
)

// 给worker分配channel
func worker(id int, c chan int) {
	for n := range c { //防止没数了还在读
		fmt.Println("wroker %d receiver %c\n", id, n)
	}
}

// 返回只能收的channel
func createrworker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c) //要创一个协程收
	return c

}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createrworker(i) //都是只能收的channel
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond) //给时间打印

}

// 缓冲区
func bufferedchannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a' //用单引号！！
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	bufferedchannel()
	channelClose()

}
