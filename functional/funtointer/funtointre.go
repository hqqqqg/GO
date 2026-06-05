// 函数实现接口
package main

import "fmt"

// --- 前面写好的三步 ---
type Greeter interface {
	Greet(name string)
}

type GreetingFunc func(name string)

func (f GreetingFunc) Greet(name string) {
	f(name)
}

// ----------------------

// 这是一个专门接收 Greeter 接口的干活函数
func sayHello(g Greeter, name string) {
	fmt.Println("准备打招呼...")
	g.Greet(name)
}

func main() {
	// 1. 我们随手写一个极其普通的匿名函数，赋给变量 myFunc
	myFunc := func(name string) {
		fmt.Printf("你好啊, %s！我是用普通函数打印的。\n", name)
	}

	// 2. 见证奇迹的时刻！
	// 我们强制把 myFunc 转换成 GreetingFunc 类型。
	// 因为 GreetingFunc 实现了 Greeter 接口，所以它可以直接传进 sayHello 里！
	sayHello(GreetingFunc(myFunc), "张三")
}
