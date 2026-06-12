// 加了个go后就会并发执行
// 非抢占式，主动交出控制权
// 多个协程可在一个或者多个线程运行
// io操作会协程切换，不io可能会霸占，top看看占用情况
// 手动交出控制权 runtime.Gosched(),很少用
// 不传i 会out of index
// go run -race goroutine.go 检测访问冲突
// 可能切换的点 i/o，select,channel ,等待锁，函数调用（maybe)，runtime.Gosched()
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { //！
			for {
				a[i]++ //一边
				runtime.Gosched()
			}
		}(i) //传进去
	}
	time.Sleep(time.Millisecond) //毫秒
	fmt.Println(a)               //一边
}
